package disabler

import (
	"fmt"
	"github.com/Layr-Labs/eigenlayer-rewards-updater/pkg/services"
	"go.uber.org/zap"
)

type Disabler struct {
	transactor services.Transactor
	logger     *zap.Logger
}

func NewDisabler(
	transactor services.Transactor,
	logger *zap.Logger,
) (*Disabler, error) {
	return &Disabler{
		transactor: transactor,
		logger:     logger,
	}, nil
}

func (d *Disabler) DisableRoot(rootIndex uint32) error {
	if rootIndex < 0 {
		return fmt.Errorf("root index must be greater than or equal to 0")
	}

	err := d.transactor.DisableRoot(rootIndex)
	if err != nil {
		d.logger.Sugar().Errorf("Failed to disable root", zap.Error(err))
		return err
	}

	return nil
}
