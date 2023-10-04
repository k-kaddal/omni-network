package keeper

import (
	"context"

	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/k-kaddal/omni/x/ethbridge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StoreState(goCtx context.Context, msg *types.MsgStoreState) (*types.MsgStoreStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetState(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "State already set")
	}
	
	state, err := FetchEthereumState(msg.Address, msg.Slot)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Fetching Ethereum state failed")
	}

	k.SetState(ctx, *state)

	return &types.MsgStoreStateResponse{}, nil
}


func FetchEthereumState(address string, storageSlot string) (*types.State, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ethRPCL := os.Getenv("ETH_RPC")

	client, err := rpc.Dial(ethRPCL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// if !isValidEthereumAddress(address) {
	// 	log.Fatalf("%v is not a valid Ethereum address", address)
	// }

	stateStorage, err := FetchEthereumStorage(client, address, storageSlot)
	if err != nil {
		log.Fatalf("Failed to fetch Storage: %v", err)
	}

	stateBalance, err := FetchEthereumBalance(client, address)
	if err != nil {
		log.Fatalf("Failed to fetch balance: %v", err)
	}

	responseJSON, err := json.Marshal(stateStorage)


	metadata := &types.Metadata{
		Address: address,
		Timestamp: "",
		BlockNumber: "91",
	}

	state := &types.State{
		Address:      address,
		Slot:  storageSlot,
		Balance: stateBalance.String(),
		Data: string(responseJSON),
		Metadata: metadata,
	}

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

