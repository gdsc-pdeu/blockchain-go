// This branch (dev) should be used for further development of this project

package main

import (
	"fmt"

	"github.com/Maharshi-Pandya/blockchain-go/blockchain"
)

// ---------------------------------------------------------------------------

func main() {

	chain := blockchain.InitBlockchain()

	chain.AddBlock("Alice sent Bob 10 ImCoins")
	chain.AddBlock("Bob sent Charlie 40 ImCoins")
	chain.AddBlock("Alice sent Charlie 2 ImCoins")

	// view the chain
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
