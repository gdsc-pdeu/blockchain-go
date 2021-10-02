package blockchain

import (
	"bytes"
	"crypto/sha256"
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

// --------------------------------------------------------------------

// temporary derive hashing function.
// TODO: add a more sophisticated function in proof.go
func (b *Block) DeriveHash() {

	// join data and prevHash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	// digest the info
	hash := sha256.Sum256(info)

	// set the hash of the block
	b.Hash = hash[:]
}
