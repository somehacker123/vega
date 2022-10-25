package api

import (
	"context"

	"code.vegaprotocol.io/vega/libs/jsonrpc"
	"github.com/mitchellh/mapstructure"
)

type ClientDisconnectWallet struct {
	sessions *Sessions
}

type ClientDisconnectWalletParams struct {
	Token string `json:"hostname"`
}

// Handle disconnect a wallet for a third-party application.
//
// It does not fail. If there's no connected wallet for the given token, nothing
// happens.
//
// The wallet resources are unloaded.
func (h *ClientDisconnectWallet) Handle(_ context.Context, rawParams jsonrpc.Params) (jsonrpc.Result, *jsonrpc.ErrorDetails) {
	params, err := validateDisconnectWalletParams(rawParams)
	if err != nil {
		return nil, invalidParams(err)
	}

	h.sessions.DisconnectWallet(params.Token)

	return nil, nil
}

func validateDisconnectWalletParams(rawParams jsonrpc.Params) (ClientDisconnectWalletParams, error) {
	if rawParams == nil {
		return ClientDisconnectWalletParams{}, ErrParamsRequired
	}

	params := ClientDisconnectWalletParams{}
	if err := mapstructure.Decode(rawParams, &params); err != nil {
		return ClientDisconnectWalletParams{}, ErrParamsDoNotMatch
	}

	if params.Token == "" {
		return ClientDisconnectWalletParams{}, ErrConnectionTokenIsRequired
	}

	return params, nil
}

func NewDisconnectWallet(sessions *Sessions) *ClientDisconnectWallet {
	return &ClientDisconnectWallet{
		sessions: sessions,
	}
}