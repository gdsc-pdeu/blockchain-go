package blockchain

import "math/big"

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

const Difficulty = 12

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
