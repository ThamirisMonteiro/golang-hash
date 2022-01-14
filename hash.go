package hash

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"
)

func New(structToBeHashed interface{}) ([]byte, error) {
	structConvertedToSliceOfBytes, err := json.Marshal(structToBeHashed)
	if err != nil {
		return nil, fmt.Errorf("unable to create new hash: %w", err)
	}
	hash := sha256.Sum256(structConvertedToSliceOfBytes)
	return hash[:], nil
}

func IsSame(firstHash, secondHash []byte) bool {
	return reflect.DeepEqual(firstHash, secondHash)
}
