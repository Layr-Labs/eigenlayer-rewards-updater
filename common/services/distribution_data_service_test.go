package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/services"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/stretchr/testify/assert"
)

var testTimestamp int64 = 1712127631

func TestGetDistributionToSubmit(t *testing.T) {
	d := createTestPaymentsTable(services.PAYMENTS_TO_SUBMIT_TABLE)

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

func TestLatestSubmittedDistribution(t *testing.T) {
	d := createTestPaymentsTable(services.LATEST_SUBMITTED_PAYMENTS_TABLE)

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

func createTestPaymentsTable(tablename string) *distribution.Distribution {
	conn.ExecSQL(fmt.Sprintf("DROP TABLE IF EXISTS localnet_local.%s;", tablename))

	conn.ExecSQL(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS localnet_local.%s (
			earner bytea,
			token bytea,
			cumulative_payment numeric,
			timestamp numeric
		);
	`, tablename))

	d := utils.GetTestDistribution()

	for accountPair := d.GetStart(); accountPair != nil; accountPair = accountPair.Next() {
		for tokenPair := accountPair.Value.Oldest(); tokenPair != nil; tokenPair = tokenPair.Next() {
			conn.ExecSQL(fmt.Sprintf(`
				INSERT INTO localnet_local.%s (earner, token, cumulative_payment, timestamp)
				VALUES ($1, $2, $3, $4);
			`, tablename), accountPair.Key.Bytes(), tokenPair.Key.Bytes(), tokenPair.Value, testTimestamp)
		}
	}

	return d
}
