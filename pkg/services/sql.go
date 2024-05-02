package services

// TODO: use correct table names
// TODO(seanmcgary): convert this to prepared statements
const getPaymentsAtTimestampQuery = `
select
  gtt.earner,
  gtt.token,
  sum(gtt.amount) as cumulative_amount
from eigenlabs_payment.gold_table_test as gtt
where snapshot <= from_unixtime(%d)
group by 1, 2
order by 1, 2
`
const GetPaymentsAtTimestampQuery = getPaymentsAtTimestampQuery

const getMaxTimestampQuery = `SELECT CAST(to_unixtime(MAX(snapshot)) AS BIGINT) FROM eigenlabs_payment.gold_table_test`
const GetMaxTimestampQuery = getMaxTimestampQuery
