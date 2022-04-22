package utils

import (
	"bytes"
	"encoding/gob"
)

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(i)
	if err != nil {
		panic(err)
	}
}
