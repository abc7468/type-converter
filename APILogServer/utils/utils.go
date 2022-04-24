package utils

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(i)
	if err != nil {
		panic(err)
	}
}

func StringToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
