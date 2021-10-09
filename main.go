// This branch (dev) should be used for further development of this project

package main

import (
	"os"

	"github.com/Maharshi-Pandya/blockchain-go/blockchain"
	"github.com/Maharshi-Pandya/blockchain-go/cli"
)

// ---------------------------------------------------------------------------

func main() {

	// to exit gracefully at last
	defer os.Exit(0)

	// since InitBlockchain opens up the DB, we close it instantly using defer
	// so as to delay it till surrounding functions return
	chain := blockchain.InitBlockchain()
	defer chain.Database.Close()

	// init cli
	cli := cli.CommandLineInterface{Chain: chain}
	cli.Run()
}
