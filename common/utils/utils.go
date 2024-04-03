package utils

import "os"

func GetEnvNetwork() string {
	return os.Getenv("ENV") + "-" + os.Getenv("NETWORK")
}
