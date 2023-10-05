package ethereumbridge_test

import (
	"testing"

	keepertest "github.com/k-kaddal/omni/testutil/keeper"
	"github.com/k-kaddal/omni/testutil/nullify"
	"github.com/k-kaddal/omni/x/ethereumbridge"
	"github.com/k-kaddal/omni/x/ethereumbridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EthStateList: []types.EthState{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthereumbridgeKeeper(t)
	ethereumbridge.InitGenesis(ctx, *k, genesisState)
	got := ethereumbridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EthStateList, got.EthStateList)
	// this line is used by starport scaffolding # genesis/test/assert
}
