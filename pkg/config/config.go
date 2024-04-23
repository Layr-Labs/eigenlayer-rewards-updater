package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type GlobalConfig struct {
	Config string `mapstructure:"config"`
	Debug  bool   `mapstructure:"debug"`
}

type Environment int

var (
	Environment_LOCAL    Environment = 0
	Environment_DEV      Environment = 1
	Environment_PRE_PROD Environment = 2
	Environment_PROD     Environment = 3
)

type UpdaterConfig struct {
	GlobalConfig
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
	case "local", "localnet":
		return Environment_LOCAL
	default:
		return Environment_DEV
	}
}

// StringEnvironmentFromEnum gets a string environment value from the enum
func StringEnvironmentFromEnum(env Environment) (string, error) {
	switch env {
	case Environment_PROD:
		return "prod", nil
	case Environment_PRE_PROD:
		return "pre-prod", nil
	case Environment_DEV:
		return "dev", nil
	case Environment_LOCAL:
		return "localnet", nil
	}
	return "", errors.New(fmt.Sprintf("String env not found for '%d'", env))
}

// NewUpdaterConfig reads config values from viper and returns
// them in a struct
func NewUpdaterConfig() *UpdaterConfig {
	updaterConfig = &UpdaterConfig{
		GlobalConfig: GlobalConfig{
			Config: viper.GetString("config"),
			Debug:  viper.GetBool("debug"),
		},
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
	envString, err := StringEnvironmentFromEnum(c.Environment)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s_%s", envString, c.Network), nil
}

func GetUpdaterConfig() *UpdaterConfig {
	return updaterConfig
}

func KebabToSnakeCase(str string) string {
	return strings.ReplaceAll(str, "-", "_")
}