package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/k-kaddal/omni/testutil/keeper"
	"github.com/k-kaddal/omni/x/ethbridge/keeper"
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EthbridgeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
