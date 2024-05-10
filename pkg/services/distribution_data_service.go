package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-proofs/pkg/distribution"
	"go.uber.org/zap"
)

var ErrNewDistributionNotCalculated = fmt.Errorf("new distribution not calculated")

type DistributionDataService interface {
	// Gets the latest calculated distribution that has not been submitted to the chain and the timestamp it which it was calculated until
	GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int32, error)

	// Gets the latest submitted distribution
	GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int32, error)
}

type DistributionDataServiceConfig struct {
	EnvNetwork string
	Logger     *zap.Logger
}

type DistributionDataServiceImpl struct {
	db         *sql.DB
	transactor Transactor
	config     *DistributionDataServiceConfig
}

func NewDistributionDataService(db *sql.DB, transactor Transactor, cfg *DistributionDataServiceConfig) DistributionDataService {
	return &DistributionDataServiceImpl{
		db:         db,
		transactor: transactor,
		config:     cfg,
	}
}

func (dds *DistributionDataServiceImpl) GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int32, error) {
	// get latest submitted timestamp from the chain
	latestSubmittedTimestamp, err := dds.transactor.CurrPaymentCalculationEndTimestamp()
	if err != nil {
		return nil, 0, err
	}

	// get the latest calculated timestamp from the database
	var timestamp int32
	err = dds.db.QueryRow(getMaxTimestampQuery).Scan(&timestamp)
	if err != nil {
		return nil, 0, err
	}

	// if the latest submitted timestamp is >= the latest calculated timestamp, return an error
	if int32(latestSubmittedTimestamp) >= timestamp {
		return nil, 0, fmt.Errorf("%w - latest submitted: %d, latest calculated: %d", ErrNewDistributionNotCalculated, latestSubmittedTimestamp, timestamp)
	}

	dds.config.Logger.Sugar().Info(
		fmt.Sprintf("Latest submitted timestamp: %d, Latest calculated timestamp: %d", latestSubmittedTimestamp, timestamp),
		zap.Int32("timestamp", timestamp),
		zap.Uint32("latestSubmittedTimestamp", latestSubmittedTimestamp),
	)

	d, err := dds.populateDistributionFromTable(ctx, timestamp)
	if err != nil {
		return nil, 0, err
	}

	return d, timestamp, err
}

func (dds *DistributionDataServiceImpl) GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int32, error) {
	// get latest submitted timestamp from the chain
	latestSubmittedTimestamp, err := dds.transactor.CurrPaymentCalculationEndTimestamp()
	if err != nil {
		return nil, 0, err
	}

	dds.config.Logger.Sugar().Debugf("Got timestamp '%d'", latestSubmittedTimestamp)

	d, err := dds.populateDistributionFromTable(ctx, int32(latestSubmittedTimestamp))
	if err != nil {
		return nil, 0, err
	}

	return d, int32(latestSubmittedTimestamp), err
}

func (dds *DistributionDataServiceImpl) populateDistributionFromTable(ctx context.Context, timestamp int32) (*distribution.Distribution, error) {
	d := distribution.NewDistribution()
	rows, err := dds.db.Query(fmt.Sprintf(GetPaymentsAtTimestampQuery, timestamp))
	if err != nil {
		return nil, err
	}

	earners := []*distribution.EarnerLine{}

	// populate the distribution
	for rows.Next() {
		var earnerString string
		var tokenString string
		var cumulativePaymentString uint64
		err := rows.Scan(&earnerString, &tokenString, &cumulativePaymentString)
		if err != nil {
			return nil, err
		}

		earners = append(earners, &distribution.EarnerLine{
			Earner:           earnerString,
			Token:            tokenString,
			CumulativeAmount: float64(cumulativePaymentString),
		})
	}
	d.LoadLines(earners)
	fmt.Printf("Distribution: %+v\n", d)
	return d, nil
}
