package calculator

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/rs/zerolog/log"
)

type DistributionDataService interface {
	// GetDistributionAtTimestamp returns the distribution of all tokens at a given timestamp
	GetDistributionAtTimestamp(timestamp *big.Int) (*distribution.Distribution, error)
	// SetDistributionAtTimestamp sets the distribution of all tokens at a given timestamp
	SetDistributionAtTimestamp(timestamp *big.Int, distributions *distribution.Distribution) error
}

type DistributionDataServiceImpl struct{}

func NewDistributionDataService() DistributionDataService {
	return &DistributionDataServiceImpl{}
}

func NewDistributionDataServiceImpl() *DistributionDataServiceImpl {
	return &DistributionDataServiceImpl{}
}

func (s *DistributionDataServiceImpl) GetDistributionAtTimestamp(timestamp *big.Int) (*distribution.Distribution, error) {
	// if the data directory doesn't exist, create it and return empty map
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		err = os.Mkdir("./data", 0755)
		if err != nil {
			return nil, err
		}
		return distribution.NewDistribution(), nil
	}
	if err != nil {
		return nil, err
	}

	// read from data/distributions_{timestamp}.json
	file, err := os.ReadFile(fmt.Sprintf("data/distribution_%d.json", timestamp))
	if err != nil {
		return nil, err
	}

	// deserialize from json
	var distribution *distribution.Distribution
	err = json.Unmarshal(file, distribution)
	if err != nil {
		return nil, err
	}

	return distribution, nil
}

func (s *DistributionDataServiceImpl) SetDistributionAtTimestamp(timestamp *big.Int, distribution *distribution.Distribution) error {
	// seralize to json and write to data/distributions_{timestamp}.json
	marshalledDistribution, err := json.Marshal(distribution)
	if err != nil {
		return err
	}

	log.Info().Msgf("marshalled distributions %s", marshalledDistribution)

	// write to file
	err = os.WriteFile(fmt.Sprintf("data/distribution_%d.json", timestamp), marshalledDistribution, 0644)
	if err != nil {
		return err
	}

	return err
}
