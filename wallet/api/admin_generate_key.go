package api

import (
	"context"
	"fmt"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	"code.vegaprotocol.io/vega/wallet/wallet"
	"github.com/mitchellh/mapstructure"
)

type GenerateKeyParams struct {
	Wallet     string            `json:"wallet"`
	Metadata   []wallet.Metadata `json:"metadata"`
	Passphrase string            `json:"passphrase"`
}

type GenerateKeyResult struct {
	PublicKey string            `json:"publicKey"`
	Algorithm wallet.Algorithm  `json:"algorithm"`
	Metadata  []wallet.Metadata `json:"metadata"`
}

type GenerateKey struct {
	walletStore WalletStore
}

// Handle generates a key of the specified wallet.
func (h *GenerateKey) Handle(ctx context.Context, rawParams jsonrpc.Params) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	params, err := validateGenerateKeyParams(rawParams)
	if err != nil {
		return nil, invalidParams(err)
	}

	if exist, err := h.walletStore.WalletExists(ctx, params.Wallet); err != nil {
		return nil, internalError(fmt.Errorf("could not verify the wallet existence: %w", err))
	} else if !exist {
		return nil, invalidParams(ErrWalletDoesNotExist)
	}

	w, err := h.walletStore.GetWallet(ctx, params.Wallet, params.Passphrase)
	if err != nil {
		return nil, internalError(fmt.Errorf("could not retrieve the wallet: %w", err))
	}

	kp, err := w.GenerateKeyPair(params.Metadata)
	if err != nil {
		return nil, internalError(fmt.Errorf("could not generate a new key: %w", err))
	}

	if err := h.walletStore.SaveWallet(ctx, w, params.Passphrase); err != nil {
		return nil, internalError(fmt.Errorf("could not save the wallet: %w", err))
	}

	return GenerateKeyResult{
		PublicKey: kp.PublicKey(),
		Algorithm: wallet.Algorithm{
			Name:    kp.AlgorithmName(),
			Version: kp.AlgorithmVersion(),
		},
		Metadata: kp.Metadata(),
	}, nil
}

func validateGenerateKeyParams(rawParams jsonrpc.Params) (GenerateKeyParams, error) {
	if rawParams == nil {
		return GenerateKeyParams{}, ErrParamsRequired
	}

	params := GenerateKeyParams{}
	if err := mapstructure.Decode(rawParams, &params); err != nil {
		return GenerateKeyParams{}, ErrParamsDoNotMatch
	}

	if params.Wallet == "" {
		return GenerateKeyParams{}, ErrWalletIsRequired
	}

	if params.Passphrase == "" {
		return GenerateKeyParams{}, ErrPassphraseIsRequired
	}

	return params, nil
}

func NewGenerateKey(
	walletStore WalletStore,
) *GenerateKey {
	return &GenerateKey{
		walletStore: walletStore,
	}
}
