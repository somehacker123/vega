package api

import (
	"context"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	walletnode "code.vegaprotocol.io/vega/wallet/api/node"
)

type ClientGetChainIDResult struct {
	ChainID string `json:"chainID"`
}

type ClientGetChainID struct {
	nodeSelector walletnode.Selector
}

func (h *ClientGetChainID) Handle(ctx context.Context) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	currentNode, err := h.nodeSelector.Node(ctx, noNodeSelectionReporting)
	if err != nil {
		return nil, NodeCommunicationError(ErrNoHealthyNodeAvailable)
	}

	lastBlockData, err := currentNode.LastBlock(ctx)
	if err != nil {
		return nil, NodeCommunicationError(ErrCouldNotGetLastBlockInformation)
	}

	return ClientGetChainIDResult{
		ChainID: lastBlockData.ChainID,
	}, nil
}

func NewGetChainID(nodeSelector walletnode.Selector) *ClientGetChainID {
	return &ClientGetChainID{
		nodeSelector: nodeSelector,
	}
}
