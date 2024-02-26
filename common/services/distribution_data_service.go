package services

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/rs/zerolog/log"
)

type DistributionDataService interface {
	// GetDistribution returns the distribution of payments for the given root
	GetDistribution(root [32]byte) (*distribution.Distribution, error)
	// SetDistribution sets the distribution of payments for the given root
	SetDistribution(root [32]byte, distribution *distribution.Distribution) error
}

type DistributionDataServiceImpl struct{}

func NewDistributionDataService() DistributionDataService {
	return &DistributionDataServiceImpl{}
}

func NewDistributionDataServiceImpl() *DistributionDataServiceImpl {
	return &DistributionDataServiceImpl{}
}

func (s *DistributionDataServiceImpl) GetDistribution(root [32]byte) (*distribution.Distribution, error) {
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
	file, err := os.ReadFile(distributionFileName(root))
	if err != nil {
		return nil, err
	}

	// deserialize from json
	distribution := &distribution.Distribution{}
	err = json.Unmarshal(file, distribution)
	if err != nil {
		return nil, err
	}

	return distribution, nil
}

func (s *DistributionDataServiceImpl) SetDistribution(root [32]byte, distribution *distribution.Distribution) error {

	// seralize to json and write to data/distributions_{root}.json
	marshalledDistribution, err := json.Marshal(distribution)
	if err != nil {
		return err
	}

	log.Info().Msgf("writing distribution to file %s", distributionFileName(root))

	// write to file
	err = os.WriteFile(distributionFileName(root), marshalledDistribution, 0644)
	if err != nil {
		return err
	}

	return err
}

func distributionFileName(root [32]byte) string {
	return fmt.Sprintf("data/distribution_%s.json", hex.EncodeToString(root[:]))
}
