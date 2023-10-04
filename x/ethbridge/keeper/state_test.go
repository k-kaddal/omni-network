package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/k-kaddal/omni/testutil/keeper"
	"github.com/k-kaddal/omni/testutil/nullify"
	"github.com/k-kaddal/omni/x/ethbridge/keeper"
	"github.com/k-kaddal/omni/x/ethbridge/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNState(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.State {
	items := make([]types.State, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetState(ctx, items[i])
	}
	return items
}

func TestStateGet(t *testing.T) {
	keeper, ctx := keepertest.EthbridgeKeeper(t)
	items := createNState(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetState(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStateRemove(t *testing.T) {
	keeper, ctx := keepertest.EthbridgeKeeper(t)
	items := createNState(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveState(ctx,
			item.Address,
		)
		_, found := keeper.GetState(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestStateGetAll(t *testing.T) {
	keeper, ctx := keepertest.EthbridgeKeeper(t)
	items := createNState(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllState(ctx)),
	)
}
