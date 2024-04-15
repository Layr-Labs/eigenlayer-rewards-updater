package main

import (
	"os"

	"database/sql"

	drv "github.com/uber/athenadriver/go"
)

const region = "us-east-1"
const outputBucket = "s3://payment-poc-mock/query-results/"

func main() {
	// Step 1. Set AWS Credential in Driver Config.
	conf, _ := drv.NewDefaultConfig(outputBucket, region, os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"))
	// Step 2. Open Connection.
	dbpool, _ := sql.Open(drv.DriverName, conf.Stringify())
	defer dbpool.Close()
	// Step 3. Query and print results
	var timestamp int64
	_ = dbpool.QueryRow("SELECT CAST(to_unixtime(MAX(calculation_timestamp)) AS BIGINT) FROM dev_devnet.cumulative_payments;").Scan(&timestamp)
	println(timestamp)
}
