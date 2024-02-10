package distribution

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// Merklizes the leaves and returns the merkle root.
func Merklize(leafs [][32]byte) ([32]byte, error) {
	if len(leafs) == 0 {
		return [32]byte{0xDE, 0xAD}, nil // todo: fix this
	}

	// todo: parallelize this
	numNodes := len(leafs)
	for numNodes > 1 {
		// if the number of leafs is odd, duplicate the last leaf
		if len(leafs)%2 == 1 {
			leafs = append(leafs, ZERO_LEAF)
			numNodes++
		}

		// combine the leafs into inner nodes
		for i := 0; i < len(leafs); i += 2 {
			leafs[i/2] = [32]byte(crypto.Keccak256(leafs[i][:], leafs[i+1][:]))
		}

		// remove the leafs that were combined
		numNodes /= 2
	}

	return leafs[0], nil
}
