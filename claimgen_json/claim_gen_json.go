package main

import (
	"encoding/json"
	"io"
	"math/big"
	"os"

	paymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	claimProver "github.com/Layr-Labs/eigenlayer-payment-updater/claimgen"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

// Solidity call compatible type structs

type IPaymentCoordinatorEarnerTreeMerkleLeafStrings struct {
	Earner          common.Address
	EarnerTokenRoot string
}

type IPaymentCoordinatorPaymentMerkleClaimStrings struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof string
	EarnerLeaf      IPaymentCoordinatorEarnerTreeMerkleLeafStrings
	LeafIndices     []uint32
	TokenTreeProofs []string
	TokenLeaves     []paymentCoordinator.IPaymentCoordinatorTokenTreeMerkleLeaf
}

type IPaymentCoordinatorTokenTreeMerkleLeafStrings struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

func main() {
	filePath, outputPath := "test_data/distribution_data.json", "test_data/data_output.json"
	rootIndex := uint32(0)

	GenerateProofFromJSONForSolidity(
		filePath,
		outputPath,
		rootIndex,
		TestAddressesJSON[0],
		[]gethcommon.Address{
			TestTokensJSON[0],
			TestTokensJSON[1],
			TestTokensJSON[2],
			TestTokensJSON[3],
		},
	)
}

func GenerateProofFromJSON(
	filePath string,
	outputPath string,
	rootIndex uint32,
	earner gethcommon.Address,
	tokens []gethcommon.Address,
) (*paymentCoordinator.IPaymentCoordinatorPaymentMerkleClaim, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	distro := distribution.NewDistribution()
	err = distro.UnmarshalJSON(byteValue)
	if err != nil {
		return nil, err
	}
	// generate the trees
	accountTree, tokenTrees, err := distro.Merklize()
	if err != nil {
		return nil, err
	}

	merkleClaim, error := claimProver.GetProof(
		distro,
		rootIndex,
		accountTree,
		tokenTrees,
		earner,
		tokens,
	)

	jsonData, err := json.Marshal(merkleClaim)
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = os.WriteFile(outputPath, jsonData, 0644) // 0644 is the file permission
	if err != nil {
		return nil, err
	}

	return merkleClaim, error
}

func GenerateProofFromJSONForSolidity(
	filePath string,
	outputPath string,
	rootIndex uint32,
	earner gethcommon.Address,
	tokens []gethcommon.Address,
) (*IPaymentCoordinatorPaymentMerkleClaimStrings, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	distro := distribution.NewDistribution()
	err = distro.UnmarshalJSON(byteValue)
	if err != nil {
		return nil, err
	}
	// generate the trees
	accountTree, tokenTrees, err := distro.Merklize()
	if err != nil {
		return nil, err
	}

	merkleClaim, error := claimProver.GetProof(
		distro,
		rootIndex,
		accountTree,
		tokenTrees,
		earner,
		tokens,
	)

	// Convert the merkle claim to a solidity compatible struct
	solidityMerkleClaim := IPaymentCoordinatorPaymentMerkleClaimStrings{
		RootIndex:       merkleClaim.RootIndex,
		EarnerIndex:     merkleClaim.EarnerIndex,
		EarnerTreeProof: utils.ConvertBytesToString(merkleClaim.EarnerTreeProof),
		EarnerLeaf: IPaymentCoordinatorEarnerTreeMerkleLeafStrings{
			Earner:          merkleClaim.EarnerLeaf.Earner,
			EarnerTokenRoot: utils.ConvertBytes32ToString(merkleClaim.EarnerLeaf.EarnerTokenRoot),
		},
		LeafIndices:     merkleClaim.LeafIndices,
		TokenTreeProofs: utils.ConvertBytesToStrings(merkleClaim.TokenTreeProofs),
		TokenLeaves:     merkleClaim.TokenLeaves,
	}

	jsonData, err := json.Marshal(solidityMerkleClaim)
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = os.WriteFile(outputPath, jsonData, 0644) // 0644 is the file permission
	if err != nil {
		return nil, err
	}

	return &solidityMerkleClaim, error
}

/// Test data from test_data/distribution_data.json

var TestAddressesJSON = []gethcommon.Address{
	gethcommon.HexToAddress("0x0D6bA28b9919CfCDb6b233469Cc5Ce30b979e08E"),
	gethcommon.HexToAddress("0x31F4155eb332C5bE0350589B8C95C65c3edddC99"),
	gethcommon.HexToAddress("0xC350b89bc87d36B6159f6451Ed9b760874ec7B17"),
	gethcommon.HexToAddress("0xF2288D736d27C1584Ebf7be5f52f9E4d47251AeE"),
	gethcommon.HexToAddress("0xaA179c1DBB8fa3Ca1c8e82919E009969Ce27b515"),
	gethcommon.HexToAddress("0xbd84d8216B8c69D38b7328f0AA27D4cf79e6309a"),
}

var TestTokensJSON = []gethcommon.Address{
	gethcommon.HexToAddress("0x1006dd1B8C3D0eF53489beD27577C75299F71473"),
	gethcommon.HexToAddress("0x11a4B85eaB283C98D27C8AE64469224D55Ed1894"),
	gethcommon.HexToAddress("0x43afFfbe0AfAcDABE9Ce7DBc4f07407a2b788A84"),
	gethcommon.HexToAddress("0x748a3eD7E6b04239150d7eBe12D7aeF3e3994A23"),
	gethcommon.HexToAddress("0xD275b23e0a5B68Ae251B0Dc4c81104CBa36E7cD6"),
	gethcommon.HexToAddress("0xec562ACb9E470DE27DcA2495950660fA9fbd85F8"),
}
