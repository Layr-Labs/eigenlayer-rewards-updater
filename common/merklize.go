package common

import "github.com/ethereum/go-ethereum/crypto"

// Merklizes the leaves and returns the merkle root.
func Merklize(leafs [][]byte) ([32]byte, error) {
	// todo: parallelize this
	numNodes := len(leafs)
	for len(leafs) > 1 {
		// if the number of leafs is odd, duplicate the last leaf
		if len(leafs)%2 == 1 {
			leafs = append(leafs, ZERO_LEAF)
			numNodes++
		}

		// combine the leafs into inner nodes
		for i := 0; i < len(leafs); i += 2 {
			leafs[i/2] = crypto.Keccak256(leafs[i], leafs[i+1])
		}

		// remove the leafs that were combined
		numNodes /= 2
	}

	var root [32]byte
	copy(root[:], leafs[0])
	return root, nil
}
