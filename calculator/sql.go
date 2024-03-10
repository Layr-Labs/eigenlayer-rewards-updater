package calculator

// todo: rename claimer to recipient
var recipientsAtTimestampQuery string = `
	SELECT DISTINCT ON (account) account, claimer
	FROM %s.claimer_set
	WHERE block_timestamp <= $1 AND encode(account, 'hex') in (%s)
	ORDER BY account, block_timestamp DESC;`

// get all stakers that have an entry in the staker_delegated table with the given operator with a block timestamp higher than the entry in the staker_undelegated table for the same staker
var stakerSetSharesAtTimestampQuery string = `
WITH staker_set_shares AS (
    SELECT DISTINCT ON (staker) staker, operator, shares
    FROM %s.staker_delegation_share
    WHERE update_block_timestamp <= $1
    AND encode(strategy, 'hex') = $2
    ORDER BY staker, update_block_timestamp DESC
)
SELECT staker, shares
FROM staker_set_shares
WHERE encode(operator, 'hex') = $3;`
