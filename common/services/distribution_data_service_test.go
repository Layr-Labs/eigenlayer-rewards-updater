package services_test

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/stretchr/testify/assert"
)

var testTimestamp int64 = 1712127631

func TestGetDistributionToSubmit(t *testing.T) {
	d := createPaymentsTable()
	createDistributionRootSubmittedsTable([]int64{testTimestamp - 1})

	dds := services.NewDistributionDataService(dbpool)

	fetchedDistribution, timestamp, err := dds.GetDistributionToSubmit(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, testTimestamp, timestamp)

	expectedAccountTree, _, err := d.Merklize()
	assert.Nil(t, err)

	fetchedAccountTree, _, err := fetchedDistribution.Merklize()
	assert.Nil(t, err)

	assert.Equal(t, expectedAccountTree.Root(), fetchedAccountTree.Root())
}

func TestGetDistributionToSubmitWhenNoNewCalculations(t *testing.T) {
	createPaymentsTable()
	createDistributionRootSubmittedsTable([]int64{testTimestamp})

	dds := services.NewDistributionDataService(dbpool)

	_, _, err := dds.GetDistributionToSubmit(context.Background())
	assert.ErrorIs(t, err, services.ErrNewDistributionNotCalculated)
}

func TestLatestSubmittedDistribution(t *testing.T) {
	d := createPaymentsTable()
	createDistributionRootSubmittedsTable([]int64{testTimestamp})

	dds := services.NewDistributionDataService(dbpool)

	fetchedDistribution, timestamp, err := dds.GetLatestSubmittedDistribution(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, testTimestamp, timestamp)

	expectedAccountTree, _, err := d.Merklize()
	assert.Nil(t, err)

	fetchedAccountTree, _, err := fetchedDistribution.Merklize()
	assert.Nil(t, err)

	assert.Equal(t, expectedAccountTree.Root(), fetchedAccountTree.Root())
}

func createPaymentsTable() *distribution.Distribution {
	conn.ExecSQL("DROP TABLE IF EXISTS localnet_local.cumulative_payments;")

	conn.ExecSQL(`
		CREATE TABLE IF NOT EXISTS localnet_local.cumulative_payments (
			earner bytea,
			token bytea,
			cumulative_payment numeric,
			timestamp numeric
		);
	`)

	d := utils.GetTestDistribution()

	for accountPair := d.GetStart(); accountPair != nil; accountPair = accountPair.Next() {
		for tokenPair := accountPair.Value.Oldest(); tokenPair != nil; tokenPair = tokenPair.Next() {
			conn.ExecSQL(`
				INSERT INTO localnet_local.cumulative_payments (earner, token, cumulative_payment, timestamp)
				VALUES ($1, $2, $3, $4);
			`, accountPair.Key.Bytes(), tokenPair.Key.Bytes(), tokenPair.Value, testTimestamp)
		}
	}

	return d
}

func createDistributionRootSubmittedsTable(timestamps []int64) {
	conn.ExecSQL(`DROP TABLE IF EXISTS localnet_local.distribution_root_submitteds;`)

	conn.ExecSQL(`
		CREATE TABLE IF NOT EXISTS localnet_local.distribution_root_submitteds (
			paymentCalculationEndTimestamp numeric
		);
	`)

	for _, timestamp := range timestamps {
		conn.ExecSQL(`
			INSERT INTO localnet_local.distribution_root_submitteds (paymentCalculationEndTimestamp)
			VALUES ($1);
		`, timestamp)
	}
}
