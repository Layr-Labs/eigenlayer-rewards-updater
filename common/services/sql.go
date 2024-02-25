package services

// gets all range payments that overlap with the given range ($1, $2)
var overlappingRangePaymentsQuery string = `
	SELECT range_payment_avs, range_payment_strategy, range_payment_token, range_payment_amount, range_payment_start_range_timestamp, range_payment_end_range_timestamp
	FROM %s.range_payment_created
	WHERE range_payment_start_range_timestamp < $2 AND range_payment_end_range_timestamp > $1 AND
	      block_timestamp >= $3 AND block_timestamp < $4;
`

var latestRootSubmissionQuery string = `
	SELECT root, payments_calculated_until_timestamp
	FROM %s.root_submitted
	ORDER BY payments_calculated_until_timestamp DESC
	LIMIT 1;
`
