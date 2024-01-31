package updater

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
)

const FINALIZATION_DEPTH = 100

type PaymentDataService interface {
	GetLatestFinalizedTimestamp(ctx context.Context) (*big.Int, error)
}

type PaymentDataServiceImpl struct {
	ChainClient *common.ChainClient
}

func NewPaymentDataService(chainClient *common.ChainClient) PaymentDataService {
	return &PaymentDataServiceImpl{
		ChainClient: chainClient,
	}
}

func (s *PaymentDataServiceImpl) GetLatestFinalizedTimestamp(ctx context.Context) (*big.Int, error) {
	latestBlockNumber, err := s.ChainClient.GetCurrentBlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	latestFinalizedBlockNumber := latestBlockNumber - FINALIZATION_DEPTH

	latestFinalizedBlock, err := s.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(latestFinalizedBlockNumber)))
	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(latestFinalizedBlock.Time)), nil
}
