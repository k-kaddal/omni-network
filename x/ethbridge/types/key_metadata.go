package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MetadataKeyPrefix is the prefix to retrieve all Metadata
	MetadataKeyPrefix = "Metadata/value/"
)

// MetadataKey returns the store key to retrieve a Metadata from the index fields
func MetadataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
