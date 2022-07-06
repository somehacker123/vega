// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package gql

import (
	"context"

	protoapi "code.vegaprotocol.io/protos/data-node/api/v1"
	types "code.vegaprotocol.io/protos/vega"
)

type myAssetResolver VegaResolverRoot

func (r *myAssetResolver) Status(ctx context.Context, obj *types.Asset) (AssetStatus, error) {
	return convertAssetStatusFromProto(obj.Status)
}

func (r *myAssetResolver) InfrastructureFeeAccount(ctx context.Context, obj *types.Asset) (*types.Account, error) {
	if len(obj.Id) <= 0 {
		return nil, ErrMissingIDOrReference
	}
	req := &protoapi.FeeInfrastructureAccountsRequest{
		Asset: obj.Id,
	}
	res, err := r.tradingDataClient.FeeInfrastructureAccounts(ctx, req)
	if err != nil {
		return nil, err
	}

	var acc *types.Account
	if len(res.Accounts) > 0 {
		acc = res.Accounts[0]
	}

	return acc, nil
}

func (r *myAssetResolver) GlobalRewardPoolAccount(ctx context.Context, obj *types.Asset) (*types.Account, error) {
	if len(obj.Id) <= 0 {
		return nil, ErrMissingIDOrReference
	}
	req := &protoapi.GlobalRewardPoolAccountsRequest{
		Asset: obj.Id,
	}
	res, err := r.tradingDataClient.GlobalRewardPoolAccounts(ctx, req)
	if err != nil {
		return nil, err
	}

	var acc *types.Account
	if len(res.Accounts) > 0 {
		acc = res.Accounts[0]
	}

	return acc, nil
}

func (r myAssetResolver) Name(ctx context.Context, obj *types.Asset) (string, error) {
	return obj.Details.Name, nil
}

func (r myAssetResolver) Symbol(ctx context.Context, obj *types.Asset) (string, error) {
	return obj.Details.Symbol, nil
}

func (r myAssetResolver) TotalSupply(ctx context.Context, obj *types.Asset) (string, error) {
	return obj.Details.TotalSupply, nil
}

func (r *myAssetResolver) Decimals(ctx context.Context, obj *types.Asset) (int, error) {
	return int(obj.Details.Decimals), nil
}

func (r *myAssetResolver) Quantum(ctx context.Context, obj *types.Asset) (string, error) {
	return obj.Details.Quantum, nil
}

func (r *myAssetResolver) Source(ctx context.Context, obj *types.Asset) (AssetSource, error) {
	return AssetSourceFromProto(obj.Details)
}

func AssetSourceFromProto(pdetails *types.AssetDetails) (AssetSource, error) {
	if pdetails == nil {
		return nil, ErrNilAssetSource
	}
	switch asimpl := pdetails.Source.(type) {
	case *types.AssetDetails_BuiltinAsset:
		return BuiltinAssetFromProto(asimpl.BuiltinAsset), nil
	case *types.AssetDetails_Erc20:
		return ERC20FromProto(asimpl.Erc20), nil
	default:
		return nil, ErrUnimplementedAssetSource
	}
}

func BuiltinAssetFromProto(ba *types.BuiltinAsset) *BuiltinAsset {
	return &BuiltinAsset{
		MaxFaucetAmountMint: ba.MaxFaucetAmountMint,
	}
}

func ERC20FromProto(ea *types.ERC20) *Erc20 {
	return &Erc20{
		ContractAddress:   ea.ContractAddress,
		LifetimeLimit:     ea.LifetimeLimit,
		WithdrawThreshold: ea.WithdrawThreshold,
	}
}
