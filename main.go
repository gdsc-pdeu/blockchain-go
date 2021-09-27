// This branch (dev) should be used for further development of this project

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block

	// Methods:
	// AddBlock(): Adds the Block to the Blockchain using data
}

// ---------------------------------------------------------------------

func (chain *Blockchain) AddBlock(data string) {
	// create the block and append
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

// --------------------------------------------------------------------

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte

	// Methods:
	// DeriveHash(): Set the Hash of current block
}

func (b *Block) DeriveHash() {

	// join data and prevHash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	// digest the info
	hash := sha256.Sum256(info)

	// set the hash of the block
	b.Hash = hash[:]
}

// ---------------------------------------------------------------------------

func CreateBlock(data string, prevHash []byte) *Block {
	// init a new block and derive its hash
	newBlock := &Block{[]byte{}, []byte(data), prevHash}
	newBlock.DeriveHash()
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

// ---------------------------------------------------------------------------

func main() {

	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	// view the chain
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
