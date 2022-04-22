package utils

import (
	"bytes"
	"encoding/gob"
)

func ToBytes(i interface{}) []byte {
	var dataBuffer bytes.Buffer
	encoder := gob.NewEncoder(&dataBuffer)
	err := encoder.Encode(i)
	if err != nil {
		panic(err)
	}
	return dataBuffer.Bytes()
}
