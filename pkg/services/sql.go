package services

// TODO: use correct table names
const getPaymentsAtTimestampQuery = `
SELECT earner, token, cumulative_amount
FROM %s.cumulative_payments
WHERE calculation_timestamp = from_unixtime(%d)
ORDER BY earner, token;
`
const GetPaymentsAtTimestampQuery = getPaymentsAtTimestampQuery

const getMaxTimestampQuery = `SELECT CAST(to_unixtime(MAX(calculation_timestamp)) AS BIGINT) FROM %s.cumulative_payments;`
const GetMaxTimestampQuery = getMaxTimestampQuery
