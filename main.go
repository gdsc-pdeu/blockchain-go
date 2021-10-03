// This branch (dev) should be used for further development of this project

package main

import (
	"fmt"
	"strconv"

	"github.com/Maharshi-Pandya/blockchain-go/blockchain"
)

// ---------------------------------------------------------------------------

func main() {

	fmt.Println("Computing hashes...")

	chain := blockchain.InitBlockchain()

	chain.AddBlock("Alice sent Bob 10 ImCoins")
	fmt.Println()
	chain.AddBlock("Bob sent Charlie 40 ImCoins")
	fmt.Println()
	chain.AddBlock("Alice sent Charlie 2 ImCoins")
	fmt.Println()

	// view the chain
	fmt.Println("\nBlockchain Generated ---")
	for _, block := range chain.Blocks {
		fmt.Printf("\nPrevious Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW validated: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
