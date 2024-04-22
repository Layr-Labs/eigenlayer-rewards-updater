package utils

import (
	"math/big"
	"os"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
)

// Utils for unit and integration tests

var ()

func SetTestEnv() {
	os.Setenv("ENV", "localnet")
	os.Setenv("NETWORK", "local")
}

func GetTestDistribution() *distribution.Distribution {
	d := distribution.NewDistribution()

	// give some addresses many tokens
	// addr1 => token_1 => 1
	// addr1 => token_2 => 2
	// ...
	// addr1 => token_n => n
	// addr2 => token_1 => 2
	// addr2 => token_2 => 3
	// ...
	// addr2 => token_n-1 => n+1
	for i := 0; i < len(TestAddresses); i++ {
		for j := 0; j < len(TestTokens)-i; j++ {
			d.Set(TestAddresses[i], TestTokens[j], big.NewInt(int64(j+i+1)))
		}
	}

	return d
}

func GetCompleteTestDistribution() *distribution.Distribution {
	d := distribution.NewDistribution()

	for i := 0; i < len(TestAddresses); i++ {
		for j := 0; j < len(TestTokens); j++ {
			d.Set(TestAddresses[i], TestTokens[j], big.NewInt(int64(j+i+2)))
		}
	}

	return d
}
