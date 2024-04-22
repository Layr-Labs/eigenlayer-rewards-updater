package config

import (
	"github.com/spf13/viper"
	"strings"
)

func KebabToSnakeCase(str string) string {
	return strings.ReplaceAll(str, "-", "_")
}

type UpdaterConfig struct {
	RPCUrl             string `mapstructure:"rpc_url"`
	PrivateKey         string `mapstructure:"private_key"`
	AWSAccessKeyId     string `mapstructure:"aws_access_key_id"`
	AWSSecretAccessKey string `mapstructure:"aws_secret_access_key"`
	AWSRegion          string `mapstructure:"aws_region"`
	S3OutputBucket     string `mapstructure:"s3_output_bucket"`
}

var updaterConfig *UpdaterConfig

// NewUpdaterConfig reads config values from viper and returns
// them in a struct
func NewUpdaterConfig() *UpdaterConfig {
	updaterConfig = &UpdaterConfig{
		RPCUrl:             viper.GetString("rpc_url"),
		PrivateKey:         viper.GetString("private_key"),
		AWSAccessKeyId:     viper.GetString("aws_access_key_id"),
		AWSSecretAccessKey: viper.GetString("aws_secret_access_key"),
		AWSRegion:          viper.GetString("aws_region"),
		S3OutputBucket:     viper.GetString("s3_output_bucket"),
	}
	return updaterConfig
}

func GetUpdaterConfig() *UpdaterConfig {
	return updaterConfig
}
