package ethbridge_test

import (
	"testing"

	keepertest "github.com/k-kaddal/omni/testutil/keeper"
	"github.com/k-kaddal/omni/testutil/nullify"
	"github.com/k-kaddal/omni/x/ethbridge"
	"github.com/k-kaddal/omni/x/ethbridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MetadataList: []types.Metadata{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthbridgeKeeper(t)
	ethbridge.InitGenesis(ctx, *k, genesisState)
	got := ethbridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MetadataList, got.MetadataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
