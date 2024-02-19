package updater

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentCalculatorDataService(t *testing.T) {
	elpds := NewPaymentsDataServiceImpl(
		dbpool,
		schemaService,
	)

	t.Run("test GetPaymentsCalculatedUntilTimestamp", func(t *testing.T) {
		createRootSubmittedTable()
		paymentsCalculatedUntilTimestamp, err := elpds.GetPaymentsCalculatedUntilTimestamp(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, paymentsCalculatedUntilTimestamp.Int64(), int64(1708385990))
	})

	// TODO: overlapping range payments test

	t.Cleanup(func() {
		conn.ExecSQL(`
			DROP TABLE IF EXISTS sgd34.root_submitted;
		`)
	})
}

func createRootSubmittedTable() {
	conn.ExecSQL(`
		CREATE TABLE IF NOT EXISTS sgd34.root_submitted (
			id bytea PRIMARY KEY,
			root bytea NOT NULL,
			payments_calculated_until_timestamp numeric NOT NULL,
			activated_after numeric NOT NULL,
			block_number numeric NOT NULL,
			block_timestamp numeric NOT NULL,
			transaction_hash bytea NOT NULL
		);
	`)

	// insert a couple rows
	conn.ExecSQL(`
		INSERT INTO sgd34.root_submitted VALUES (
			decode('1234', 'hex'),
			decode('0000000000000000000000000000000000000000000000000000000000000001', 'hex'),
			1708285990,
			2708285990,
			10559231,
			1708293276,
			decode('1234567890123456789012345678901234567890123456789012345678901234', 'hex')
		);
	`)

	conn.ExecSQL(`
		INSERT INTO sgd34.root_submitted VALUES (
			decode('5678', 'hex'),
			decode('0000000000000000000000000000000000000000000000000000000000000002', 'hex'),
			1708385990,
			3708385990,
			10569231,
			1708493276,
			decode('5678901234567890123456789012345678901234567890123456789012345678', 'hex')
		);
	`)
}
