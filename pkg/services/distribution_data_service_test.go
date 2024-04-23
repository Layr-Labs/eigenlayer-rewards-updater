package services_test

import (
	"context"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/mocks"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/config"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/distribution"
	services2 "github.com/Layr-Labs/eigenlayer-payment-updater/pkg/services"
	utils2 "github.com/Layr-Labs/eigenlayer-payment-updater/pkg/utils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var testTimestamp int64 = 1712127631

func TestGetDistributionToSubmit(t *testing.T) {
	cfg := config.UpdaterConfig{
		Environment: config.Environment_LOCAL,
		Network:     "local",
	}

	networkEnv, err := cfg.GetEnvNetwork()
	if err != nil {
		t.Fatalf("Failed to get EnvNetwork")
	}

	// return test timestamp from chain
	mockTransactor := &mocks.Transactor{}
	mockTransactor.On("CurrPaymentCalculationEndTimestamp").Return(uint64(testTimestamp), nil)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create rows
	d, rows := getDistributionAndPaymentRows()

	// return testTimestamp + 1 from db, so we've calculated a new distribution
	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(services2.GetMaxTimestampQuery, networkEnv))).WillReturnRows(getMaxTimestampRows(testTimestamp + 1))
	// return the distribution rows
	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(services2.GetPaymentsAtTimestampQuery, networkEnv, testTimestamp+1))).WillReturnRows(rows)

	dds := services2.NewDistributionDataService(db, mockTransactor, &services2.DistributionDataServiceConfig{
		EnvNetwork: networkEnv,
	})

	fetchedDistribution, timestamp, err := dds.GetDistributionToSubmit(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, testTimestamp+1, timestamp)

	expectedAccountTree, _, err := d.Merklize()
	assert.Nil(t, err)

	fetchedAccountTree, _, err := fetchedDistribution.Merklize()
	assert.Nil(t, err)

	assert.Equal(t, expectedAccountTree.Root(), fetchedAccountTree.Root())
}

func TestGetDistributionToSubmitWhenNoNewCalculations(t *testing.T) {
	cfg := config.UpdaterConfig{
		Environment: config.Environment_LOCAL,
		Network:     "local",
	}

	networkEnv, err := cfg.GetEnvNetwork()
	if err != nil {
		t.Fatalf("Failed to get EnvNetwork")
	}

	mockTransactor := &mocks.Transactor{}
	mockTransactor.On("CurrPaymentCalculationEndTimestamp").Return(uint64(testTimestamp), nil)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// return testTimestamp from db, so we haven't calculated a new distribution
	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(services2.GetMaxTimestampQuery, networkEnv))).WillReturnRows(getMaxTimestampRows(testTimestamp))

	dds := services2.NewDistributionDataService(db, mockTransactor, &services2.DistributionDataServiceConfig{
		EnvNetwork: networkEnv,
	})

	_, _, err = dds.GetDistributionToSubmit(context.Background())
	assert.ErrorIs(t, err, services2.ErrNewDistributionNotCalculated)
}

func TestLatestSubmittedDistribution(t *testing.T) {
	cfg := config.UpdaterConfig{
		Environment: config.Environment_LOCAL,
		Network:     "local",
	}

	networkEnv, err := cfg.GetEnvNetwork()
	if err != nil {
		t.Fatalf("Failed to get EnvNetwork")
	}

	mockTransactor := &mocks.Transactor{}
	mockTransactor.On("CurrPaymentCalculationEndTimestamp").Return(uint64(testTimestamp), nil)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create rows
	d, rows := getDistributionAndPaymentRows()

	// return the distribution at testTimestamp from db
	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf(services2.GetPaymentsAtTimestampQuery, networkEnv, testTimestamp))).WillReturnRows(rows)

	dds := services2.NewDistributionDataService(db, mockTransactor, &services2.DistributionDataServiceConfig{
		EnvNetwork: networkEnv,
	})

	fetchedDistribution, timestamp, err := dds.GetLatestSubmittedDistribution(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, testTimestamp, timestamp)

	expectedAccountTree, _, err := d.Merklize()
	assert.Nil(t, err)

	fetchedAccountTree, _, err := fetchedDistribution.Merklize()
	assert.Nil(t, err)

	assert.Equal(t, expectedAccountTree.Root(), fetchedAccountTree.Root())
}

func getDistributionAndPaymentRows() (*distribution.Distribution, *sqlmock.Rows) {
	d := utils2.GetTestDistribution()

	rows := sqlmock.NewRows([]string{"eaner", "token", "culumative_payment"})

	for accountPair := d.GetStart(); accountPair != nil; accountPair = accountPair.Next() {
		for tokenPair := accountPair.Value.Oldest(); tokenPair != nil; tokenPair = tokenPair.Next() {
			rows.AddRow(accountPair.Key.String(), tokenPair.Key.String(), decimal.NewFromBigInt(tokenPair.Value.Int, 0))
		}
	}

	return d, rows
}

func getMaxTimestampRows(timestamp int64) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"_col0"}).AddRow(timestamp)
}
