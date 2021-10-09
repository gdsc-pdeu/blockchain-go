package blockchain

import (
	"fmt"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

// define the db options
const (
	dbPath = "./tmp/blocks"
)

// chain of blocks
type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

// iterator for the chain
type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// individual block
type Block struct {
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Nonce     int
	TimeStamp int64
}

// ---------------------------------------------------------------------

func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// 1) View the db to find the hash of the last block
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.ValueCopy(nil)

		return err
	})
	Handle(err)

	// 2) Update the db by creating a new block and setting the key to appropriate hash
	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.ToByteSlice())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		// set the in memory lastHash
		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err)
}

// ---------------------------------------------------------------------------

func CreateBlock(data string, prevHash []byte) *Block {
	// init a new block and derive its hash
	newBlock := &Block{[]byte{}, []byte(data), prevHash, 0, time.Now().UnixNano()}

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

// init the database and store the genesis block
func InitBlockchain() *Blockchain {
	// to store the hash of last block in memory
	var lastHash []byte

	// open the database
	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	Handle(err)

	// check if the key "lh" exists already
	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("\n::-> No blockchain found. Creating one...")
			genesis := Genesis()
			fmt.Println("\n::-> Found the genesis block...")

			err = txn.Set(genesis.Hash, genesis.ToByteSlice())
			Handle(err)

			// set the lh key to last block's hash
			err = txn.Set([]byte("lh"), genesis.Hash)

			// set the in memory lastHash
			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.ValueCopy(nil)
			return err
		}
	})
	Handle(err)

	// init the blockchain
	blockchain := Blockchain{lastHash, db}
	return &blockchain
}

// initialises the iterator
func (chain *Blockchain) InitIterator() *BlockchainIterator {
	bci := BlockchainIterator{chain.LastHash, chain.Database}
	return &bci
}

// Get the next block from the blockchain
func (bci *BlockchainIterator) Next() *Block {
	var block *Block

	// get the entry corresponding to lastHash from DB
	err := bci.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(bci.CurrentHash)
		Handle(err)

		// since we stored the block in serialized form
		serialBlock, err := item.ValueCopy(nil)
		block = FromByteSlice(serialBlock)

		return err
	})
	Handle(err)

	// set the current hash of iterator to previous hash of the block
	bci.CurrentHash = block.PrevHash

	return block
}
