package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/k-kaddal/omni/x/ethereumbridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EthStateAll(c context.Context, req *types.QueryAllEthStateRequest) (*types.QueryAllEthStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ethStates []types.EthState
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ethStateStore := prefix.NewStore(store, types.KeyPrefix(types.EthStateKeyPrefix))

	pageRes, err := query.Paginate(ethStateStore, req.Pagination, func(key []byte, value []byte) error {
		var ethState types.EthState
		if err := k.cdc.Unmarshal(value, &ethState); err != nil {
			return err
		}

		ethStates = append(ethStates, ethState)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEthStateResponse{EthState: ethStates, Pagination: pageRes}, nil
}

func (k Keeper) EthState(c context.Context, req *types.QueryGetEthStateRequest) (*types.QueryGetEthStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEthState(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEthStateResponse{EthState: val}, nil
}
