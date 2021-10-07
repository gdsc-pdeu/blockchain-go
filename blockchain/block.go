package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// chain of blocks
type Blockchain struct {
	Blocks []*Block

	// Methods:
	// AddBlock(): Adds the Block to the Blockchain using data
}

// individual block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int

	// Methods:
	// DeriveHash(): Set the Hash of current block
}

// ---------------------------------------------------------------------

func (chain *Blockchain) AddBlock(data string) {
	// create the block and append
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

// ---------------------------------------------------------------------------

func CreateBlock(data string, prevHash []byte) *Block {
	// init a new block and derive its hash
	newBlock := &Block{[]byte{}, []byte(data), prevHash, 0}

	// init a new pow to mine the block
	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.Mine()

	// set the hash and nonce for the block
	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce

	return newBlock
}

func Genesis() *Block {
	// create the genesis block
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	bls := []*Block{Genesis()}
	ch := &Blockchain{bls}
	return ch
}

// TODO: create serializing and deserializing functions to convert the Block
// into a byte array and vice versa for storing in badger database
func (block *Block) ToByteSlice() []byte {
	var buf bytes.Buffer

	// gob encoder
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(block)

	Handle(err)

	return buf.Bytes()
}

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
