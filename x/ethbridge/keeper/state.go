package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/k-kaddal/omni/x/ethbridge/types"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"

	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
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


func (k Keeper) FetchEthereumState(address string, storageSlot string) (*types.State, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ethRPCL := os.Getenv("ETH_RPC")

	client, err := rpc.Dial(ethRPCL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if !k.isValidEthereumAddress(address) {
		log.Fatalf("%v is not a valid Ethereum address", address)
	}

	stateStorage, err := FetchEthereumStorage(client, address, storageSlot)
	if err != nil {
		log.Fatalf("Failed to fetch Storage: %v", err)
	}

	stateBalance, err := FetchEthereumBalance(client, address)
	if err != nil {
		log.Fatalf("Failed to fetch balance: %v", err)
	}

	responseJSON, err := json.Marshal(stateStorage)


	// TODO: Fetch the block number
	var blockNumber string
	err = client.Call(&blockNumber, "eth_blockNumber")
	if err != nil {
		return nil, err
	}

	metadata := &types.Metadata{
		Address: address,
		Timestamp: time.Now().String(),
		BlockNumber: blockNumber,
	}

	state := &types.State{
		Address:      address+"/"+storageSlot,
		Slot:  storageSlot,
		Balance: stateBalance.String(),
		Data: string(responseJSON),
		Metadata: metadata,
	}

	fmt.Print(state)

	return state, nil
}

func FetchEthereumBalance(client *rpc.Client, address string) (*big.Int, error) {
	var stateBalance string
	err := client.Call(&stateBalance, "eth_getBalance", address, "latest")
	if err != nil {
		log.Fatalf("Error eth_getBalance at address: %v", address)
		return nil, err
	}

	balanceWei := new(big.Int)
	balanceWei, ok := balanceWei.SetString(stateBalance, 0)
	if !ok {
		return nil, fmt.Errorf("Failed to convert balance to wei")
	}

	return balanceWei, nil
}

func FetchEthereumStorage(client *rpc.Client, address string, storageSlot string) (string, error) {
	var stateStorage string
	err := client.Call(&stateStorage, "eth_getStorageAt", address, storageSlot, "latest")

	if err != nil {
		log.Fatalf("Error eth_getStorageAt at address: %v", address)
	}

	return stateStorage, nil
}
