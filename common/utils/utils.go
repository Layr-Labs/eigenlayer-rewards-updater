package utils

import "os"

func GetEnvNetwork() string {
	return os.Getenv("ENV") + "_" + os.Getenv("NETWORK")
}
