package distribution

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// Merklizes the leaves and returns the merkle root.
func SimpleMerklize(leafs [][32]byte) ([32]byte, error) {
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

// Merklizes the leaves and returns the merkle root.
func MerklizeAndCacheLayers(leafs [][32]byte) ([32]byte, [][][32]byte, error) {
	if len(leafs) == 0 {
		return [32]byte{0xDE, 0xAD}, nil, nil // todo: fix this
	}

	// todo: parallelize this
	layers := make([][][32]byte, 0)
	layers = append(layers, leafs) // add the leaf layer

	layer := layers[0]
	for len(layer) > 1 {
		// if the number of leafs is odd, duplicate the last leaf
		if len(layer)%2 == 1 {
			layer = append(layer, ZERO_LEAF)
		}

		newLayer := make([][32]byte, len(layer)/2)
		// combine the leafs into inner nodes
		for i := 0; i < len(layer); i += 2 {
			newLayer[i/2] = [32]byte(crypto.Keccak256(layer[i][:], layer[i+1][:]))
		}

		// add the new layer
		layers = append(layers, newLayer)
		layer = newLayer
	}

	// the first element of the last remaining layer is the root
	return layer[0], layers, nil
}
