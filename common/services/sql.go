package services

const PAYMENTS_TO_SUBMIT_TABLE = "cumulative_payments_to_submit"
const LATEST_SUBMITTED_PAYMENTS_TABLE = "latest_submitted_cumulative_payments"

// TODO: use correct table names
const lockTableForReadsQuery = `LOCK TABLE %s.%s IN SHARE MODE;`

const getAllPaymentsBalancesQuery = `SELECT earner, token, cumulative_payment FROM %s.%s ORDER BY earner, token;`

const getTimestampQuery = `SELECT timestamp FROM %s.%s LIMIT 1;`
