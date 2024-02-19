package updater

var paymentsCalculatedUntilQuery string = `
	SELECT payments_calculated_until_timestamp
	FROM %s.root_submitted
	ORDER BY payments_calculated_until_timestamp DESC
	LIMIT 1;
`
