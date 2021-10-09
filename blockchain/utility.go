package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
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

// Serialize a Block to a byte slice
func (block *Block) ToByteSlice() []byte {
	var buf bytes.Buffer

	// gob encoder
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(block)

	Handle(err)

	return buf.Bytes()
}

// Deserialize a byte slice to get the corresponding block
func FromByteSlice(serial []byte) *Block {
	var block Block

	// gob decoder
	decoder := gob.NewDecoder(bytes.NewReader(serial))
	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Error handling
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
