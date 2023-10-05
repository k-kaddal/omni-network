package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/k-kaddal/omni/x/ethereumbridge/types"
)

func (k msgServer) StoreState(goCtx context.Context, msg *types.MsgStoreState) (*types.MsgStoreStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetEthState(
		ctx,
		msg.Address+"-"+msg.Slot,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "State already set")
	}

	state, err := k.FetchEthereumState(msg.Address, msg.Slot)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Fetching Ethereum state failed")
	}
	
	err = k.VerifyEthereumState(state)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Ethereum State not valid")
	}
	
	k.SetEthState(ctx, *state)

	return &types.MsgStoreStateResponse{}, nil
}
