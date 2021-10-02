package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

// This file contains the utility functions
func IntToByteSlice(num int64) []byte {
	// initialize a new bytes buffer
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	// buffer now contain bytes of the num
	return buf.Bytes()
}
