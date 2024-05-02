package updater_test

import (
	"context"
	"github.com/Layr-Labs/eigenlayer-payment-updater/internal/logger"
	"github.com/Layr-Labs/eigenlayer-payment-updater/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/updater"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testTimestamp int32 = 1712127631

func TestUpdaterUpdate(t *testing.T) {
	logger, _ := logger.NewLogger(&logger.LoggerConfig{Debug: true})

	mockTransactor := &mocks.Transactor{}
	mockDistributionDataService := &mocks.DistributionDataService{}

	d := utils.GetTestDistribution()
	accountTree, _, _ := d.Merklize()
	rootBytes := accountTree.Root()
	var root [32]byte
	copy(root[:], rootBytes)

	mockDistributionDataService.On("GetDistributionToSubmit", mock.Anything).Return(d, testTimestamp, nil)
	mockTransactor.On("SubmitRoot", mock.Anything, root, big.NewInt(int64(testTimestamp))).Return(nil)

	updater, err := updater.NewUpdater(mockTransactor, mockDistributionDataService, logger)
	assert.Nil(t, err)

	_, err = updater.Update(context.Background())

	assert.Nil(t, err)
}
