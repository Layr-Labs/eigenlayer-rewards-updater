package updater_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/Layr-Labs/eigenlayer-payment-updater/updater"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testUpdaterIntervalSeconds = 10
var testTimestamp int64 = 1712127631

func TestUpdaterUpdate(t *testing.T) {
	mockTransactor := &mocks.Transactor{}
	mockDistributionDataService := &mocks.DistributionDataService{}

	d := utils.GetTestDistribution()
	accountTree, _, _ := d.Merklize()
	rootBytes := accountTree.Root()
	var root [32]byte
	copy(root[:], rootBytes)

	mockDistributionDataService.On("GetDistributionToSubmit", mock.Anything).Return(d, testTimestamp, nil)
	mockTransactor.On("SubmitRoot", mock.Anything, root, big.NewInt(testTimestamp)).Return(nil)

	updater, err := updater.NewUpdater(testUpdaterIntervalSeconds, mockTransactor, mockDistributionDataService)
	assert.Nil(t, err)

	err = updater.Update(context.Background())

	assert.Nil(t, err)
}
