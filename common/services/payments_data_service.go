package services

import (
	"context"
	"fmt"
	"math/big"

	contractIPaymentCoordinator "github.com/Layr-Labs/eigenlayer-payment-updater/bindings/IPaymentCoordinator"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common"
	"github.com/Layr-Labs/eigenlayer-payment-updater/common/utils"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type RangePayment struct {
	*contractIPaymentCoordinator.IPaymentCoordinatorRangePayment
	Avs gethcommon.Address
}

type PaymentsDataService interface {
	// GetLatestRootSubmission returns the latest root and timestamp until which payments have been calculated
	GetLatestRootSubmission(ctx context.Context) ([32]byte, *big.Int, error)
	// GetRangePaymentsWithOverlappingRange returns all range payments that overlap with the given range that were created within the given creation range
	GetRangePaymentsWithOverlappingRange(rangeStartTimestamp, rangeEndTimestamp, creationStartTimestamp, creationEndTimestamp *big.Int) ([]*RangePayment, error)
}

type PaymentsDataServiceImpl struct {
	dbpool        *pgxpool.Pool
	schemaService *common.SubgraphSchemaService
}

func NewPaymentsDataService(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
) PaymentsDataService {
	return &PaymentsDataServiceImpl{
		dbpool:        dbpool,
		schemaService: schemaService,
	}
}

func NewPaymentsDataServiceImpl(
	dbpool *pgxpool.Pool,
	schemaService *common.SubgraphSchemaService,
) *PaymentsDataServiceImpl {
	return &PaymentsDataServiceImpl{
		dbpool:        dbpool,
		schemaService: schemaService,
	}
}

func (s *PaymentsDataServiceImpl) GetLatestRootSubmission(ctx context.Context) ([32]byte, *big.Int, error) {
	schemaID, err := s.schemaService.GetSubgraphSchema(ctx, utils.SUBGRAPH_CLAIMING_MANAGER)
	if err != nil {
		return [32]byte{}, nil, err
	}

	formattedQuery := fmt.Sprintf(latestRootSubmissionQuery, schemaID)
	row := s.dbpool.QueryRow(ctx, formattedQuery)

	var resRootBytes []byte
	var resTimestampDecimal decimal.Decimal
	err = row.Scan(
		&resRootBytes,
		&resTimestampDecimal,
	)
	if err != nil {
		return [32]byte{}, nil, err
	}

	fixedResRootBytes := [32]byte{}
	copy(fixedResRootBytes[:], resRootBytes)

	return fixedResRootBytes, resTimestampDecimal.BigInt(), nil
}

func (s *PaymentsDataServiceImpl) GetRangePaymentsWithOverlappingRange(rangeStartTimestamp, rangeEndTimestamp, creationStartTimestamp, creationEndTimestamp *big.Int) ([]*RangePayment, error) {
	schemaID, err := s.schemaService.GetSubgraphSchema(context.Background(), utils.SUBGRAPH_PAYMENT_COORDINATOR)
	if err != nil {
		return nil, err
	}

	formattedQuery := fmt.Sprintf(overlappingRangePaymentsQuery, schemaID)

	log.Info().Msgf("executing query: %s, %d, %d, %d, %d", formattedQuery, rangeStartTimestamp, rangeEndTimestamp, creationStartTimestamp, creationEndTimestamp)

	rows, err := s.dbpool.Query(context.Background(), formattedQuery, rangeStartTimestamp, rangeEndTimestamp, creationStartTimestamp, creationEndTimestamp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rangePayments []*RangePayment
	for rows.Next() {
		// var (
		// 	rangePaymentAVSBytes                   []byte
		// 	rangePaymentStrategyBytes              []byte
		// 	rangePaymentTokenBytes                 []byte
		// 	rangePaymentAmountDecimal              decimal.Decimal
		// 	rangePaymentStartRangeTimestampDecimal decimal.Decimal
		// 	rangePaymentEndRangeTimestampDecimal   decimal.Decimal
		// )

		// err := rows.Scan(
		// 	&rangePaymentAVSBytes,
		// 	&rangePaymentStrategyBytes,
		// 	&rangePaymentTokenBytes,
		// 	&rangePaymentAmountDecimal,
		// 	&rangePaymentStartRangeTimestampDecimal,
		// 	&rangePaymentEndRangeTimestampDecimal,
		// )
		// if err != nil {
		// 	return nil, err
		// }

		// rangePayment := &contractIPaymentCoordinator.IPaymentCoordinatorRangePayment{
		// 	Avs:                 gethcommon.HexToAddress(hex.EncodeToString(rangePaymentAVSBytes)),
		// 	Strategy:            gethcommon.HexToAddress(hex.EncodeToString(rangePaymentStrategyBytes)),
		// 	Token:               gethcommon.HexToAddress(hex.EncodeToString(rangePaymentTokenBytes)),
		// 	Amount:              rangePaymentAmountDecimal.BigInt(),
		// 	StartRangeTimestamp: rangePaymentStartRangeTimestampDecimal.BigInt(),
		// 	EndRangeTimestamp:   rangePaymentEndRangeTimestampDecimal.BigInt(),
		// }

		// rangePayments = append(rangePayments, rangePayment)
	}
	return rangePayments, nil
}
