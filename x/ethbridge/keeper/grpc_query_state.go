package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/k-kaddal/omni/x/ethbridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StateAll(c context.Context, req *types.QueryAllStateRequest) (*types.QueryAllStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var states []types.State
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	stateStore := prefix.NewStore(store, types.KeyPrefix(types.StateKeyPrefix))

	pageRes, err := query.Paginate(stateStore, req.Pagination, func(key []byte, value []byte) error {
		var state types.State
		if err := k.cdc.Unmarshal(value, &state); err != nil {
			return err
		}

		states = append(states, state)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStateResponse{State: states, Pagination: pageRes}, nil
}

func (k Keeper) State(c context.Context, req *types.QueryGetStateRequest) (*types.QueryGetStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetState(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStateResponse{State: val}, nil
}
