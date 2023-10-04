package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

// SetState set a specific state in the store from its index
func (k Keeper) SetState(ctx sdk.Context, state types.State) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateKeyPrefix))
	b := k.cdc.MustMarshal(&state)
	store.Set(types.StateKey(
		state.Address,
	), b)
}

// GetState returns a state from its index
func (k Keeper) GetState(
	ctx sdk.Context,
	address string,

) (val types.State, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateKeyPrefix))

	b := store.Get(types.StateKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveState removes a state from the store
func (k Keeper) RemoveState(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateKeyPrefix))
	store.Delete(types.StateKey(
		address,
	))
}

// GetAllState returns all state
func (k Keeper) GetAllState(ctx sdk.Context) (list []types.State) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.State
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
