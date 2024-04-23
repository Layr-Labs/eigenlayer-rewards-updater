package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

func KebabToSnakeCase(str string) string {
	return strings.ReplaceAll(str, "-", "_")
}

type Environment int

var (
	Environment_DEV      Environment = 0
	Environment_PRE_PROD Environment = 1
	Environment_PROD     Environment = 2
)

type UpdaterConfig struct {
	Environment               Environment `mapstructure:"environment"`
	Network                   string      `mapstructure:"network"`
	RPCUrl                    string      `mapstructure:"rpc_url"`
	PrivateKey                string      `mapstructure:"private_key"`
	AWSAccessKeyId            string      `mapstructure:"aws_access_key_id"`
	AWSSecretAccessKey        string      `mapstructure:"aws_secret_access_key"`
	AWSRegion                 string      `mapstructure:"aws_region"`
	S3OutputBucket            string      `mapstructure:"s3_output_bucket"`
	PaymentCoordinatorAddress string      `mapstructure:"payment_coordinator_address"`
}

var updaterConfig *UpdaterConfig

// parseEnvironment normalizes environment names to an enum value
func parseEnvironment(env string) Environment {
	switch env {
	case "pre-prod", "preprod":
		return Environment_PRE_PROD
	case "prod", "production":
		return Environment_PROD
	default:
		return Environment_DEV
	}
}

// stringEnvironmentFromEnum gets a string environment value from the enum
func stringEnvironmentFromEnum(env Environment) (string, error) {
	switch env {
	case Environment_PROD:
		return "prod", nil
	case Environment_PRE_PROD:
		return "pre-prod", nil
	case Environment_DEV:
		return "dev", nil
	}
	return "", errors.New(fmt.Sprintf("String env not found for '%d'", env))
}

// NewUpdaterConfig reads config values from viper and returns
// them in a struct
func NewUpdaterConfig() *UpdaterConfig {
	updaterConfig = &UpdaterConfig{
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		AWSAccessKeyId:            viper.GetString("aws_access_key_id"),
		AWSSecretAccessKey:        viper.GetString("aws_secret_access_key"),
		AWSRegion:                 viper.GetString("aws_region"),
		S3OutputBucket:            viper.GetString("s3_output_bucket"),
		PaymentCoordinatorAddress: viper.GetString("payment_coordinator_address"),
	}
	return updaterConfig
}

// GetEnvNetwork returns a string concatenation of "{environment}_{network}"
func (c *UpdaterConfig) GetEnvNetwork() (string, error) {
	envString, err := stringEnvironmentFromEnum(c.Environment)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s_%s", envString, c.Network), nil
}

func GetUpdaterConfig() *UpdaterConfig {
	return updaterConfig
}
