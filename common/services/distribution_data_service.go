package services

import (
	"context"
	"fmt"

	"github.com/Layr-Labs/eigenlayer-payment-updater/common/distribution"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

type DistributionDataService interface {
	// Gets the latest calculated distribution that has not been submitted to the chain and the timestamp it which it was calculated until
	GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int64, error)

	// Gets the latest submitted distribution
	GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int64, error)
}

type DistributionDataServiceImpl struct {
	dbpool *pgxpool.Pool
}

func NewDistributionDataService(dbpool *pgxpool.Pool) DistributionDataService {
	return &DistributionDataServiceImpl{
		dbpool: dbpool,
	}
}

func NewDistributionDataServiceImpl(dbpool *pgxpool.Pool) *DistributionDataServiceImpl {
	return &DistributionDataServiceImpl{
		dbpool: dbpool,
	}
}

func (dds *DistributionDataServiceImpl) GetDistributionToSubmit(ctx context.Context) (*distribution.Distribution, int64, error) {
	return dds.populateDistributionFromTable(ctx, PAYMENTS_TO_SUBMIT_TABLE)
}

func (dds *DistributionDataServiceImpl) GetLatestSubmittedDistribution(ctx context.Context) (*distribution.Distribution, int64, error) {
	return dds.populateDistributionFromTable(ctx, LATEST_SUBMITTED_PAYMENTS_TABLE)
}

func (dds *DistributionDataServiceImpl) populateDistributionFromTable(ctx context.Context, table string) (*distribution.Distribution, int64, error) {
	d := distribution.NewDistribution()
	rows, err := dds.dbpool.Query(ctx, fmt.Sprintf(getAllPaymentsBalancesQuery, utils.GetEnvNetwork(), table))
	if err != nil {
		return nil, 0, err
	}

	// populate the distribution
	for rows.Next() {
		var earnerBytes []byte
		var tokenBytes []byte
		var cumulativePaymentDecimal decimal.Decimal
		err := rows.Scan(&earnerBytes, &tokenBytes, &cumulativePaymentDecimal)
		if err != nil {
			return nil, 0, err
		}

		earner := gethcommon.BytesToAddress(earnerBytes)
		token := gethcommon.BytesToAddress(tokenBytes)

		d.Set(earner, token, cumulativePaymentDecimal.BigInt())
	}

	// get the timestamp
	var timestamp int64
	err = dds.dbpool.QueryRow(ctx, fmt.Sprintf(getTimestampQuery, utils.GetEnvNetwork(), table)).Scan(&timestamp)
	if err != nil {
		return nil, 0, err
	}

	return d, timestamp, nil
}
