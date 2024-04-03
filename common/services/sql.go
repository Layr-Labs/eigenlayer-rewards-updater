package services

const PAYMENTS_TO_SUBMIT_TABLE = "%s.cumulative_payments_to_submit"
const LATEST_SUBMITTED_PAYMENTS_TABLE = "%s.latest_submitted_cumulative_payments"

// TODO: use correct table names
const lockTableForReadsQuery = `LOCK TABLE %s IN SHARE MODE;`

const getAllPaymentsBalancesQuery = `SELECT earner, token, cumulative_payment FROM %s`

const getTimestampQuery = `SELECT timestamp FROM %s LIMIT 1;`
