package keeper

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/k-kaddal/omni/x/ethereumbridge/types"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
)

// SetEthState set a specific ethState in the store from its index
func (k Keeper) SetEthState(ctx sdk.Context, ethState types.EthState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKeyPrefix))
	b := k.cdc.MustMarshal(&ethState)
	store.Set(types.EthStateKey(
		ethState.Index,
	), b)
}

// GetEthState returns a ethState from its index
func (k Keeper) GetEthState(
	ctx sdk.Context,
	index string,

) (val types.EthState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKeyPrefix))

	b := store.Get(types.EthStateKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEthState removes a ethState from the store
func (k Keeper) RemoveEthState(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKeyPrefix))
	store.Delete(types.EthStateKey(
		index,
	))
}

// GetAllEthState returns all ethState
func (k Keeper) GetAllEthState(ctx sdk.Context) (list []types.EthState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthStateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EthState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}


func (k Keeper) FetchEthereumState(address string, storageSlot string) (*types.EthState, error) {

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

	stateStorage, err := k.FetchEthereumStorage(client, address, storageSlot)
	if err != nil {
		log.Fatalf("Failed to fetch Storage: %v", err)
	}

	stateBalance, err := k.FetchEthereumBalance(client, address)
	if err != nil {
		log.Fatalf("Failed to fetch balance: %v", err)
	}

	blockNumber, err := k.FetchBlockNumer(client)
	if err != nil {
		log.Fatalf("Failed to fetch blocknumber: %v", err)
	}

	responseJSON, err := json.Marshal(stateStorage)


	state := &types.EthState{
		Index:  address + "-" + storageSlot,
		Address:  address,
		Slot:     storageSlot,
		Balance:  stateBalance.String(),
		Storage: string(responseJSON),
		Blocknumber: blockNumber,
	}

	return state, nil
}

func (k Keeper) FetchEthereumBalance(client *rpc.Client, address string) (*big.Int, error) {
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

func (k Keeper) FetchEthereumStorage(client *rpc.Client, address string, storageSlot string) (string, error) {
	var stateStorage string
	err := client.Call(&stateStorage, "eth_getStorageAt", address, storageSlot, "latest")

	if err != nil {
		log.Fatalf("Error eth_getStorageAt at address: %v", address)
	}

	return stateStorage, nil
}

func (k Keeper)  FetchBlockNumer(client *rpc.Client) (string, error) {
		var blockNumber string
		err := client.Call(&blockNumber, "eth_blockNumber")
		if err != nil {
			return "", err
		}
	return blockNumber, nil
}
