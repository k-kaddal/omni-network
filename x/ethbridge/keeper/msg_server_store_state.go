package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

func (k msgServer) StoreState(goCtx context.Context, msg *types.MsgStoreState) (*types.MsgStoreStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgStoreStateResponse{}, nil
}
