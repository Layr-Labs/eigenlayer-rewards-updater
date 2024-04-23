package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Layr-Labs/eigenlayer-payment-updater/pkg/distribution"
	"go.uber.org/zap"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

var ErrNewDistributionNotCalculated = fmt.Errorf("new distribution not calculated")

type DistributionDataService interface {
	// Gets the latest calculated distribution that has not been submitted to the chain and the timestamp it which it was calculated until
	GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int64, error)

	// Gets the latest submitted distribution
	GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int64, error)
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

func (dds *DistributionDataServiceImpl) GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int64, error) {
	// get latest submitted timestamp from the chain
	latestSubmittedTimestamp, err := dds.transactor.CurrPaymentCalculationEndTimestamp()
	if err != nil {
		return nil, 0, err
	}

	// get the latest calculated timestamp from the database
	var timestamp int64
	err = dds.db.QueryRow(fmt.Sprintf(getMaxTimestampQuery, dds.config.EnvNetwork)).Scan(&timestamp)
	if err != nil {
		return nil, 0, err
	}

	// if the latest submitted timestamp is >= the latest calculated timestamp, return an error
	if int64(latestSubmittedTimestamp) >= timestamp {
		return nil, 0, fmt.Errorf("%w - latest submitted: %d, latest calculated: %d", ErrNewDistributionNotCalculated, latestSubmittedTimestamp, timestamp)
	}

	dds.config.Logger.Sugar().Info(
		fmt.Sprintf("Latest submitted timestamp: %d, Latest calculated timestamp: %d", latestSubmittedTimestamp, timestamp),
		zap.Int64("timestamp", timestamp),
		zap.Uint64("latestSubmittedTimestamp", latestSubmittedTimestamp),
	)

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
	rows, err := dds.db.Query(fmt.Sprintf(GetPaymentsAtTimestampQuery, dds.config.EnvNetwork, timestamp))
	if err != nil {
		return nil, err
	}

	// populate the distribution
	for rows.Next() {
		var earnerString string
		var tokenString string
		var cumulativePaymentString string
		err := rows.Scan(&earnerString, &tokenString, &cumulativePaymentString)
		if err != nil {
			return nil, err
		}

		earner := gethcommon.HexToAddress(earnerString)
		token := gethcommon.HexToAddress(tokenString)

		cumulativePayment, ok := new(big.Int).SetString(cumulativePaymentString, 10)
		if !ok {
			// todo return error
			dds.config.Logger.Sugar().Error(
				fmt.Sprintf("not a valid big integer: %s", cumulativePaymentString),
				zap.String("cumulativePaymentString", cumulativePaymentString),
			)
			cumulativePayment = big.NewInt(0)

			//return nil, fmt.Errorf("not a valid big integer: %s", cumulativePaymentString)
		}

		d.Set(earner, token, cumulativePayment)
	}

	return d, nil
}
