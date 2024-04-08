package services

const PAYMENTS_TO_SUBMIT_TABLE = "cumulative_payments_to_submit"
const LATEST_SUBMITTED_PAYMENTS_TABLE = "latest_submitted_cumulative_payments"

// TODO: use correct table names
const getPaymentsAtTimestamp = `
SELECT earner, token, cumulative_payment
FROM %s.cumulative_payments
WHERE timestamp = %d
ORDER BY earner, token;
`

const getMaxTimestampQuery = `SELECT MAX(timestamp) FROM %s.cumulative_payments;`

const getPaymentsCalculatedUntilTimestamp = `SELECT MAX(paymentCalculationEndTimestamp) FROM %s.distribution_root_submitteds;`
