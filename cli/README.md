# CLI for Blockchain-Go

This package attempts to develop a command line interface for the blockchain. This requires that the blockchain is persistent across different runs of this program and thus it is stored persistently in [BadgerDB](https://github.com/dgraph-io/badger) on disk using Key-Value pairs.

Currently in this CLI, the commands are:

- **Add Block Command**: Adds a block in the blockchain with some data in it (`flag` for this command should be `-block` or `-b`)

```shell
$ go build main.go
$ ./main add -block "<any data for the block>"
```

- **Print Chain Command**: Prints the whole blockchain, from the last block to the genesis block

```shell
$ go build main.go
$ ./main print
```



Any commands to enhance this CLI can be added further :thought_balloon: