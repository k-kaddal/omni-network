package ethbridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/k-kaddal/omni/x/ethbridge/keeper"
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the metadata
	for _, elem := range genState.MetadataList {
		k.SetMetadata(ctx, elem)
	}
	// Set all the state
	for _, elem := range genState.StateList {
		k.SetState(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MetadataList = k.GetAllMetadata(ctx)
	genesis.StateList = k.GetAllState(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
