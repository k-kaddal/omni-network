package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/k-kaddal/omni/testutil/keeper"
	"github.com/k-kaddal/omni/testutil/nullify"
	"github.com/k-kaddal/omni/x/ethereumbridge/keeper"
	"github.com/k-kaddal/omni/x/ethereumbridge/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEthState(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EthState {
	items := make([]types.EthState, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEthState(ctx, items[i])
	}
	return items
}

func TestEthStateGet(t *testing.T) {
	keeper, ctx := keepertest.EthereumbridgeKeeper(t)
	items := createNEthState(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEthState(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEthStateRemove(t *testing.T) {
	keeper, ctx := keepertest.EthereumbridgeKeeper(t)
	items := createNEthState(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEthState(ctx,
			item.Index,
		)
		_, found := keeper.GetEthState(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEthStateGetAll(t *testing.T) {
	keeper, ctx := keepertest.EthereumbridgeKeeper(t)
	items := createNEthState(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEthState(ctx)),
	)
}
