package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EthStateKeyPrefix is the prefix to retrieve all EthState
	EthStateKeyPrefix = "EthState/value/"
)

// EthStateKey returns the store key to retrieve a EthState from the index fields
func EthStateKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
