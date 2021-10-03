package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

/*
	Proof of Work (PoW):

	It is a consensus for the blockchain, which everyone will agree on based on
	amount of "work" that has been put behind each block. In practice, the requirements
	of the computational work is finding a hash with some amount of leading zeroes.

	Steps:

	- Collect data from the block (Data, PrevHash, Nonce)

	- Init a counter (nonce i.e. number used once) starting at 0

	- Create a hash of data plus the counter

	- Check if the hash meets the requirements
*/

const Difficulty = 18

// PoW struct
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	// we set the target number (hashes must be less than that)
	target := big.NewInt(1)

	// left shift "target" (i.e. 1) according to the difficulty and considering
	// the fact that SHA-256 generates a 256 bit hash
	target.Lsh(target, uint(256-Difficulty))

	pow := ProofOfWork{b, target}
	return &pow
}

// "Work" to mine the blocks and derive the hashes of a block
// Returns: Nonce and Derived Hash of the block
func (pow *ProofOfWork) Mine() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	// loop through every int
	for nonce < math.MaxInt64 {
		data := pow.InitData(int64(nonce))
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		// compare the hash with the target
		// big.Int.Cmp returns:
		// -1 if x < y
		// 0 if x == y
		// 1 if  x > y
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Once mined, validate the block
func (pow *ProofOfWork) Validate() bool {
	// init the data,
	// calculate hash from the derived nonce and compare

	var intHash big.Int

	data := pow.InitData(int64(pow.Block.Nonce))
	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// Join blockData, PrevHash, Nonce and Difficulty
func (pow *ProofOfWork) InitData(nonce int64) []byte {
	// join block data, prevhash, nonce and difficulty
	data := bytes.Join([][]byte{
		[]byte(pow.Block.Data),
		pow.Block.PrevHash,
		IntToByteSlice(int64(nonce)),
		IntToByteSlice(int64(Difficulty)),
	}, []byte{})

	return data
}
