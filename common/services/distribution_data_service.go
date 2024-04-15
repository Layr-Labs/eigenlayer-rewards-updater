package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/shopspring/decimal"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

var ErrNewDistributionNotCalculated = fmt.Errorf("new distribution not calculated")

type DistributionDataService interface {
	// Gets the latest calculated distribution that has not been submitted to the chain and the timestamp it which it was calculated until
	GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int64, error)

	// Gets the latest submitted distribution
	GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int64, error)
}

type DistributionDataServiceImpl struct {
	db         *sql.DB
	transactor Transactor
}

func NewDistributionDataService(db *sql.DB, transactor Transactor) DistributionDataService {
	return &DistributionDataServiceImpl{
		db:         db,
		transactor: transactor,
	}
}

func (dds *DistributionDataServiceImpl) GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int64, error) {
	// get latest submitted timestamp from the chain
	latestSubmittedTimestamp, err := dds.transactor.CurrPaymentCalculationEndTimestamp()
	if err != nil {
		return nil, 0, err
	}

	// get the latest calculated timestamp from the database
	var timestamp int64
	err = dds.db.QueryRow(fmt.Sprintf(getMaxTimestampQuery, utils.GetEnvNetwork())).Scan(&timestamp)
	if err != nil {
		return nil, 0, err
	}

	// if the latest submitted timestamp is >= the latest calculated timestamp, return an error
	if int64(latestSubmittedTimestamp) >= timestamp {
		return nil, 0, fmt.Errorf("%w - latest submitted: %d, latest calculated: %d", ErrNewDistributionNotCalculated, latestSubmittedTimestamp, timestamp)
	}

	d, err := dds.populateDistributionFromTable(ctx, timestamp)
	if err != nil {
		return nil, 0, err
	}

	return d, timestamp, err
}

func (dds *DistributionDataServiceImpl) GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int64, error) {
	// get latest submitted timestamp from the chain
	latestSubmittedTimestamp, err := dds.transactor.CurrPaymentCalculationEndTimestamp()
	if err != nil {
		return nil, 0, err
	}

	d, err := dds.populateDistributionFromTable(ctx, int64(latestSubmittedTimestamp))
	if err != nil {
		return nil, 0, err
	}

	return d, int64(latestSubmittedTimestamp), err
}

func (dds *DistributionDataServiceImpl) populateDistributionFromTable(ctx context.Context, timestamp int64) (*distribution.Distribution, error) {
	d := distribution.NewDistribution()
	rows, err := dds.db.Query(fmt.Sprintf(GetPaymentsAtTimestampQuery, utils.GetEnvNetwork(), timestamp))
	if err != nil {
		return nil, err
	}

	// populate the distribution
	for rows.Next() {
		var earnerString string
		var tokenString string
		var cumulativePaymentDecimal decimal.Decimal
		err := rows.Scan(&earnerString, &tokenString, &cumulativePaymentDecimal)
		if err != nil {
			return nil, err
		}

		earner := gethcommon.HexToAddress(earnerString)
		token := gethcommon.HexToAddress(tokenString)

		d.Set(earner, token, cumulativePaymentDecimal.BigInt())
	}

	return d, nil
}
