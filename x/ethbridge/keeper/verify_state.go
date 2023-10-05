package keeper

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/k-kaddal/omni/x/ethbridge/types"
)

func (k Keeper) VerifyEthereumState(state *types.State) error {
	if state == nil || state.Address == "" || state.Balance == "" {
		return errors.New("Ethereum state is invalid")
	}

	if !k.isValidEthereumAddress(state.Address) {
		return errors.New("Not a valid Ethereum Address")
	}

	if !k.isPositiveNumericValue(state.Balance) {
		return errors.New("Balance is not a positive numeric value")
	}

	return nil
}

func (k Keeper) isValidEthereumAddress(address string) bool {
	trimedAddress := strings.TrimPrefix(address, "0x")

	if len(trimedAddress) != 40 {
		return false
	}

	decodedAddress := common.HexToAddress(trimedAddress)
	isValid := common.IsHexAddress(trimedAddress) && strings.ToLower(decodedAddress.Hex()) == strings.ToLower("0x"+trimedAddress)
	return isValid
}

func (k Keeper) isPositiveNumericValue(balance string) bool {
	var balanceInt big.Int
	_, ok := balanceInt.SetString(balance, 10)

	if !ok {
		errors.New("Invalid string format")
		return false
	}

	return balanceInt.Sign() >= 0
}
