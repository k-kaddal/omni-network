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

func createNMetadata(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Metadata {
	items := make([]types.Metadata, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMetadata(ctx, items[i])
	}
	return items
}

func TestMetadataGet(t *testing.T) {
	keeper, ctx := keepertest.EthbridgeKeeper(t)
	items := createNMetadata(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMetadata(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMetadataRemove(t *testing.T) {
	keeper, ctx := keepertest.EthbridgeKeeper(t)
	items := createNMetadata(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMetadata(ctx,
			item.Index,
		)
		_, found := keeper.GetMetadata(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestMetadataGetAll(t *testing.T) {
	keeper, ctx := keepertest.EthbridgeKeeper(t)
	items := createNMetadata(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMetadata(ctx)),
	)
}
