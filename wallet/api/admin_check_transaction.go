package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"code.vegaprotocol.io/vega/commands"
	vgcrypto "code.vegaprotocol.io/vega/libs/crypto"
	"code.vegaprotocol.io/vega/libs/jsonrpc"
	commandspb "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	walletpb "code.vegaprotocol.io/vega/protos/vega/wallet/v1"
	"code.vegaprotocol.io/vega/wallet/api/node"
	wcommands "code.vegaprotocol.io/vega/wallet/commands"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/mitchellh/mapstructure"
)

type AdminCheckTransactionParams struct {
	Wallet                 string        `json:"wallet"`
	PublicKey              string        `json:"publicKey"`
	Network                string        `json:"network"`
	NodeAddress            string        `json:"nodeAddress"`
	Retries                uint64        `json:"retries"`
	MaximumRequestDuration time.Duration `json:"maximumRequestDuration"`
	Transaction            interface{}   `json:"transaction"`
}

type ParsedAdminCheckTransactionParams struct {
	Wallet                 string
	PublicKey              string
	Network                string
	NodeAddress            string
	Retries                uint64
	RawTransaction         string
	MaximumRequestDuration time.Duration
}

type AdminCheckTransactionResult struct {
	ReceivedAt         time.Time               `json:"receivedAt"`
	SentAt             time.Time               `json:"sentAt"`
	Transaction        *commandspb.Transaction `json:"transaction"`
	Node               AdminNodeInfoResult     `json:"node"`
	EncodedTransaction string                  `json:"encodedTransaction"`
}

type AdminCheckTransaction struct {
	walletStore         WalletStore
	networkStore        NetworkStore
	nodeSelectorBuilder NodeSelectorBuilder
}

func (h *AdminCheckTransaction) Handle(ctx context.Context, rawParams jsonrpc.Params) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	params, err := validateAdminCheckTransactionParams(rawParams)
	if err != nil {
		return nil, InvalidParams(err)
	}

	receivedAt := time.Now()

	if exist, err := h.walletStore.WalletExists(ctx, params.Wallet); err != nil {
		return nil, InternalError(fmt.Errorf("could not verify the wallet exists: %w", err))
	} else if !exist {
		return nil, InvalidParams(ErrWalletDoesNotExist)
	}

	alreadyUnlocked, err := h.walletStore.IsWalletAlreadyUnlocked(ctx, params.Wallet)
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not verify whether the wallet is already unlock or not: %w", err))
	}
	if !alreadyUnlocked {
		return nil, RequestNotPermittedError(ErrWalletIsLocked)
	}

	w, err := h.walletStore.GetWallet(ctx, params.Wallet)
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not retrieve the wallet: %w", err))
	}

	request := &walletpb.SubmitTransactionRequest{}
	if err := jsonpb.Unmarshal(strings.NewReader(params.RawTransaction), request); err != nil {
		return nil, InvalidParams(fmt.Errorf("the transaction does not use a valid Vega command: %w", err))
	}

	request.PubKey = params.PublicKey
	request.Propagate = true
	if errs := wcommands.CheckSubmitTransactionRequest(request); !errs.Empty() {
		return nil, InvalidParams(errs)
	}

	currentNode, errDetails := h.getNode(ctx, params)
	if errDetails != nil {
		return nil, errDetails
	}

	lastBlockData, errDetails := h.getLastBlockDataFromNetwork(ctx, currentNode)
	if errDetails != nil {
		return nil, errDetails
	}

	marshaledInputData, err := wcommands.ToMarshaledInputData(request, lastBlockData.BlockHeight)
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not marshal the input data: %w", err))
	}

	signature, err := w.SignTx(params.PublicKey, commands.BundleInputDataForSigning(marshaledInputData, lastBlockData.ChainID))
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not check the transaction: %w", err))
	}

	// Build the transaction.
	tx := commands.NewTransaction(params.PublicKey, marshaledInputData, &commandspb.Signature{
		Value:   signature.Value,
		Algo:    signature.Algo,
		Version: signature.Version,
	})

	// Generate the proof of work for the transaction.
	txID := vgcrypto.RandomHash()
	powNonce, _, err := vgcrypto.PoW(lastBlockData.BlockHash, txID, uint(lastBlockData.ProofOfWorkDifficulty), lastBlockData.ProofOfWorkHashFunction)
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not compute the proof-of-work: %w", err))
	}
	tx.Pow = &commandspb.ProofOfWork{
		Nonce: powNonce,
		Tid:   txID,
	}

	sentAt := time.Now()
	if err := currentNode.CheckTransaction(ctx, tx); err != nil {
		return nil, NetworkErrorFromTransactionError(err)
	}

	rawTx, err := proto.Marshal(tx)
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not marshal the transaction: %w", err))
	}

	return AdminCheckTransactionResult{
		ReceivedAt:  receivedAt,
		SentAt:      sentAt,
		Transaction: tx,
		Node: AdminNodeInfoResult{
			Host: currentNode.Host(),
		},
		EncodedTransaction: base64.StdEncoding.EncodeToString(rawTx),
	}, nil
}

func (h *AdminCheckTransaction) getNode(ctx context.Context, params ParsedAdminCheckTransactionParams) (node.Node, *jsonrpc.ErrorDetails) {
	hosts := []string{params.NodeAddress}
	if len(params.Network) != 0 {
		exists, err := h.networkStore.NetworkExists(params.Network)
		if err != nil {
			return nil, InternalError(fmt.Errorf("could not determine if the network exists: %w", err))
		} else if !exists {
			return nil, InvalidParams(ErrNetworkDoesNotExist)
		}

		n, err := h.networkStore.GetNetwork(params.Network)
		if err != nil {
			return nil, InternalError(fmt.Errorf("could not retrieve the network configuration: %w", err))
		}

		if err := n.EnsureCanConnectGRPCNode(); err != nil {
			return nil, InvalidParams(ErrNetworkConfigurationDoesNotHaveGRPCNodes)
		}
		hosts = n.API.GRPC.Hosts
	}

	nodeSelector, err := h.nodeSelectorBuilder(hosts, params.Retries, params.MaximumRequestDuration)
	if err != nil {
		return nil, InternalError(fmt.Errorf("could not initialize the node selector: %w", err))
	}

	currentNode, err := nodeSelector.Node(ctx, noNodeSelectionReporting)
	if err != nil {
		return nil, NodeCommunicationError(ErrNoHealthyNodeAvailable)
	}

	return currentNode, nil
}

func (h *AdminCheckTransaction) getLastBlockDataFromNetwork(ctx context.Context, node node.Node) (*AdminLastBlockData, *jsonrpc.ErrorDetails) {
	lastBlock, err := node.LastBlock(ctx)
	if err != nil {
		return nil, NodeCommunicationError(ErrCouldNotGetLastBlockInformation)
	}

	if lastBlock.ChainID == "" {
		return nil, NodeCommunicationError(ErrCouldNotGetChainIDFromNode)
	}

	return &AdminLastBlockData{
		BlockHash:               lastBlock.BlockHash,
		ChainID:                 lastBlock.ChainID,
		BlockHeight:             lastBlock.BlockHeight,
		ProofOfWorkHashFunction: lastBlock.ProofOfWorkHashFunction,
		ProofOfWorkDifficulty:   lastBlock.ProofOfWorkDifficulty,
	}, nil
}

func NewAdminCheckTransaction(
	walletStore WalletStore, networkStore NetworkStore, nodeSelectorBuilder NodeSelectorBuilder,
) *AdminCheckTransaction {
	return &AdminCheckTransaction{
		walletStore:         walletStore,
		networkStore:        networkStore,
		nodeSelectorBuilder: nodeSelectorBuilder,
	}
}

func validateAdminCheckTransactionParams(rawParams jsonrpc.Params) (ParsedAdminCheckTransactionParams, error) {
	if rawParams == nil {
		return ParsedAdminCheckTransactionParams{}, ErrParamsRequired
	}

	params := AdminCheckTransactionParams{}
	if err := mapstructure.Decode(rawParams, &params); err != nil {
		return ParsedAdminCheckTransactionParams{}, ErrParamsDoNotMatch
	}

	if params.Wallet == "" {
		return ParsedAdminCheckTransactionParams{}, ErrWalletIsRequired
	}

	if params.PublicKey == "" {
		return ParsedAdminCheckTransactionParams{}, ErrPublicKeyIsRequired
	}

	if params.Network == "" && params.NodeAddress == "" {
		return ParsedAdminCheckTransactionParams{}, ErrNetworkOrNodeAddressIsRequired
	}

	if params.Network != "" && params.NodeAddress != "" {
		return ParsedAdminCheckTransactionParams{}, ErrSpecifyingNetworkAndNodeAddressIsNotSupported
	}

	if params.Transaction == nil || params.Transaction == "" {
		return ParsedAdminCheckTransactionParams{}, ErrTransactionIsRequired
	}

	tx, err := json.Marshal(params.Transaction)
	if err != nil {
		return ParsedAdminCheckTransactionParams{}, ErrTransactionIsNotValidJSON
	}

	return ParsedAdminCheckTransactionParams{
		Wallet:         params.Wallet,
		PublicKey:      params.PublicKey,
		RawTransaction: string(tx),
		Network:        params.Network,
		NodeAddress:    params.NodeAddress,
		Retries:        params.Retries,
	}, nil
}
