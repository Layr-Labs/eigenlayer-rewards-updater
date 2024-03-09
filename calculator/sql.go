package calculator

// todo: rename claimer to recipient
var recipientsAtTimestampQuery string = `
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
