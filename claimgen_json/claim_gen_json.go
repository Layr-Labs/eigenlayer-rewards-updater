package main

import (
	"encoding/json"
	claimProver "github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/claimgen"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/paymentCoordinator"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"io"
	"os"
)

// TODO: Update this to take CLI arguments to generate proofs
func main() {
	/*currentUnixTime := time.Now().Unix()
	inputPath := "claimgen_json/test_data/distribution_data3.json"
	outputPath := fmt.Sprintf("claimgen_json/test_data/data_output_%d.json", currentUnixTime)
	rootIndex := uint32(0)
	earnerIndex := uint32(3)

	_, err := GenerateProofFromJSONForSolidity(
		inputPath,
		outputPath,
		rootIndex,
		TestAddressesJSON[earnerIndex],
		[]gethcommon.Address{
			TestTokensJSON[0],
			TestTokensJSON[1],
			TestTokensJSON[2],
			TestTokensJSON[3],
			TestTokensJSON[4],
			TestTokensJSON[5],
		},
		true,
	)
	if err != nil {
		log.Fatalln(err)
	}*/
	/*addrString := gethcommon.HexToAddress("0x2222aac0c980cc029624b7ff55b88bc6f63c538f")
	tokens := []gethcommon.Address{
		gethcommon.HexToAddress("0x94373a4919b3240d86ea41593d5eba789fef3848"),
	}

	distro := distribution.NewDistribution()*/

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

	distro, err := distribution.NewDistributionWithData(byteValue)
	if err != nil {
		return nil, err
	}

	cg := claimProver.NewClaimgen(distro)

	_, merkleClaim, err := cg.GenerateClaimProofForEarner(
		earner,
		tokens,
		rootIndex,
	)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(merkleClaim)
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = os.WriteFile(outputPath, jsonData, 0644) // 0644 is the file permission
	if err != nil {
		return nil, err
	}

	return merkleClaim, nil
}

func GenerateProofFromJSONForSolidity(
	inputTreeFilePath string,
	outputPath string,
	rootIndex uint32,
	earner gethcommon.Address,
	tokens []gethcommon.Address,
	prettyJson bool,
) (*claimProver.IPaymentCoordinatorPaymentMerkleClaimStrings, error) {
	jsonFile, err := os.Open(inputTreeFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	distro, err := distribution.NewDistributionWithData(byteValue)
	if err != nil {
		return nil, err
	}

	cg := claimProver.NewClaimgen(distro)

	accountTree, merkleClaim, err := cg.GenerateClaimProofForEarner(
		earner,
		tokens,
		rootIndex,
	)
	if err != nil {
		return nil, err
	}

	solidityMerkleClaim := claimProver.FormatProofForSolidity(accountTree.Root(), merkleClaim)

	var jsonData []byte
	if prettyJson {
		jsonData, err = json.MarshalIndent(solidityMerkleClaim, "", "  ")
	} else {
		jsonData, err = json.Marshal(solidityMerkleClaim)
	}

	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = os.WriteFile(outputPath, jsonData, 0644) // 0644 is the file permission
	if err != nil {
		return nil, err
	}

	return solidityMerkleClaim, nil
}

// Test data from test_data/distribution_data.json
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
