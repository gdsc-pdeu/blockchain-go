// This branch (dev) should be used for further development of this project

package main

import (
	"fmt"

	"github.com/Maharshi-Pandya/blockchain-go/blockchain"
)

// ---------------------------------------------------------------------------

func main() {

	chain := blockchain.InitBlockchain()

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
