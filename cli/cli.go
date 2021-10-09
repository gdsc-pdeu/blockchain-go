package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/Maharshi-Pandya/blockchain-go/blockchain"
)

// reference to the chain
type CommandLineInterface struct {
	Chain *blockchain.Blockchain
}

// Print Usage
func (cli *CommandLineInterface) PrintUsage() {
	fmt.Println("--- CLI for Blockchain-Go ---")
	fmt.Println("📝 Usage:")
	fmt.Println("\n`add -block <SOME_DATA>` will add the block with the data into the blockchain")
	fmt.Println("`print` will print the entire blockchain")
}

// validation of arguments
func (cli *CommandLineInterface) validateArgs() {
	if len(os.Args) < 2 {
		fmt.Println("❌ add or print subcommand is required")
		runtime.Goexit()
	}
}

// command : Add Block
func (cli *CommandLineInterface) AddBlockCmd(data string) {
	cli.Chain.AddBlock(data)
	fmt.Println("\nAdded block into the chain ✔️")
}

// command : print the chain
func (cli *CommandLineInterface) PrintChainCmd() {
	iter := cli.Chain.InitIterator()

	// loop through the chain
	for {
		nextBlock := iter.Next()

		fmt.Printf("\nPrevious Hash: %x\n", nextBlock.PrevHash)
		fmt.Printf("Data: %s\n", nextBlock.Data)
		fmt.Printf("Hash: %x\n", nextBlock.Hash)
		fmt.Printf("Timestamp: %v\n", nextBlock.TimeStamp)

		pow := blockchain.NewProofOfWork(nextBlock)
		fmt.Printf("PoW validated: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		// check if we reach the genesis block
		if len(nextBlock.PrevHash) == 0 {
			break
		}
	}
}

// Run the CLI
func (cli *CommandLineInterface) Run() {
	cli.validateArgs()

	addBlockSubCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainSubCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockCmdData := addBlockSubCmd.String("block", "", "Data of the Block")

	// switch on the subcommand
	// FlagSet.Parse() requires a set of arguments to parse
	// arguments start from os.Args[2:]
	switch os.Args[1] {
	case "add":
		err := addBlockSubCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "print":
		err := printChainSubCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.PrintUsage()
		runtime.Goexit()
	}

	// check which command was parsed
	if addBlockSubCmd.Parsed() {
		// check for empty values
		if *addBlockCmdData == "" {
			addBlockSubCmd.Usage()
			runtime.Goexit()
		}
		cli.AddBlockCmd(*addBlockCmdData)
	}
	if printChainSubCmd.Parsed() {
		// print the blockchain
		cli.PrintChainCmd()
	}
}
