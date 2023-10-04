package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

// SetMetadata set a specific metadata in the store from its index
func (k Keeper) SetMetadata(ctx sdk.Context, metadata types.Metadata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataKeyPrefix))
	b := k.cdc.MustMarshal(&metadata)
	store.Set(types.MetadataKey(
		metadata.Index,
	), b)
}

// GetMetadata returns a metadata from its index
func (k Keeper) GetMetadata(
	ctx sdk.Context,
	index string,

) (val types.Metadata, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataKeyPrefix))

	b := store.Get(types.MetadataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMetadata removes a metadata from the store
func (k Keeper) RemoveMetadata(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataKeyPrefix))
	store.Delete(types.MetadataKey(
		index,
	))
}

// GetAllMetadata returns all metadata
func (k Keeper) GetAllMetadata(ctx sdk.Context) (list []types.Metadata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MetadataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Metadata
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
