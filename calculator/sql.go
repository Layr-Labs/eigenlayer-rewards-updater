package calculator

// gets all range payments that overlap with the given range ($1, $2)
var overlappingRangePaymentsQuery string = `
	SELECT range_payment_avs, range_payment_strategy, range_payment_token, range_payment_amount, range_payment_start_range_timestamp, range_payment_end_range_timestamp
	FROM %s.range_payment_created
	WHERE range_payment_start_range_timestamp < $2 AND range_payment_end_range_timestamp > $1
	LIMIT 1;
`

var paymentsCalculatedUntilQuery string = `
	SELECT payments_calculated_until_timestamp
	FROM %s.root_submitted
	ORDER BY payments_calculated_until_timestamp DESC
	LIMIT 1;
`

var commissionAtTimestampQuery string = `
	SELECT DISTINCT ON (operator) operator, commission_bips
	FROM %s.commission_set
	WHERE block_timestamp <= $1 AND encode(avs, 'hex') = $2 AND encode(operator, 'hex') in (%s)
	ORDER BY operator, block_timestamp DESC;`

var claimersAtTimestampQuery string = `
	SELECT DISTINCT ON (account) account, claimer
	FROM %s.claimer_set
	WHERE block_timestamp <= $1 AND encode(account, 'hex') in (%s)
	ORDER BY account, block_timestamp DESC;`

// get all stakers that have an entry in the staker_delegated table with the given operator with a block timestamp higher than the entry in the staker_undelegated table for the same staker
var stakerSetAtTimestampQuery string = `
WITH latest_undelegations AS (
    SELECT
        DISTINCT ON (staker) staker,
        block_timestamp AS undelegation_timestamp
    FROM
        %s.staker_undelegated
    WHERE
        encode(operator, 'hex') = $1
		AND block_timestamp <= $2
    ORDER BY
        staker,
        block_timestamp DESC
), latest_delegations AS (
    SELECT
        DISTINCT ON (staker) staker,
        block_timestamp AS delegation_timestamp
    FROM
        %s.staker_delegated
    WHERE
        encode(operator, 'hex') = $1
		AND block_timestamp <= $2
    ORDER BY
        staker,
        block_timestamp DESC
)
SELECT
    ld.staker
FROM
    latest_delegations ld
LEFT JOIN
    latest_undelegations lu ON ld.staker = lu.staker
WHERE
    lu.staker IS NULL OR ld.delegation_timestamp > lu.undelegation_timestamp;`
