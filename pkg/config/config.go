package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type GlobalConfig struct {
	Config      string `mapstructure:"config"`
	Debug       bool   `mapstructure:"debug"`
	DDStatsdUrl string `mapstructure:"dd_statsd_url"`
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
	RewardsCoordinatorAddress string      `mapstructure:"rewards_coordinator_address"`
	ProofStoreBaseUrl         string      `mapstructure:"proof_store_base_url"`
}

type DistributionConfig struct {
	GlobalConfig
	Environment               Environment `mapstructure:"environment"`
	Network                   string      `mapstructure:"network"`
	RPCUrl                    string      `mapstructure:"rpc_url"`
	PrivateKey                string      `mapstructure:"private_key"`
	RewardsCoordinatorAddress string      `mapstructure:"rewards_coordinator_address"`
	Output                    string      `mapstructure:"output"`
	ProofStoreBaseUrl         string      `mapstructure:"proof_store_base_url"`
}
type ClaimConfig struct {
	GlobalConfig
	Environment               Environment `mapstructure:"environment"`
	Network                   string      `mapstructure:"network"`
	RPCUrl                    string      `mapstructure:"rpc_url"`
	PrivateKey                string      `mapstructure:"private_key"`
	RewardsCoordinatorAddress string      `mapstructure:"rewards_coordinator_address"`
	Output                    string      `mapstructure:"output"`
	EarnerAddress             string      `mapstructure:"earner_address"`
	Tokens                    []string    `mapstructure:"tokens"`
	ProofStoreBaseUrl         string      `mapstructure:"proof_store_base_url"`
	ClaimTimestamp            string      `mapstructure:"claim_timestamp"`
}

var updaterConfig *UpdaterConfig
var distributionConfig *DistributionConfig
var claimConfig *ClaimConfig

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
		return "preprod", nil
	case Environment_DEV:
		return "dev", nil
	case Environment_LOCAL:
		return "local", nil
	}
	return "", errors.New(fmt.Sprintf("String env not found for '%d'", env))
}

// NewUpdaterConfig reads config values from viper and returns
// them in a struct
func NewUpdaterConfig() *UpdaterConfig {
	updaterConfig = &UpdaterConfig{
		GlobalConfig: GlobalConfig{
			Config:      viper.GetString("config"),
			Debug:       viper.GetBool("debug"),
			DDStatsdUrl: viper.GetString("dd_statsd_url"),
		},
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		RewardsCoordinatorAddress: viper.GetString("rewards_coordinator_address"),
		ProofStoreBaseUrl:         viper.GetString("proof_store_base_url"),
	}
	return updaterConfig
}

func NewDistributionConfig() *DistributionConfig {
	distributionConfig = &DistributionConfig{
		GlobalConfig: GlobalConfig{
			Config: viper.GetString("config"),
			Debug:  viper.GetBool("debug"),
		},
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		RewardsCoordinatorAddress: viper.GetString("rewards_coordinator_address"),
		Output:                    viper.GetString("output"),
		ProofStoreBaseUrl:         viper.GetString("proof_store_base_url"),
	}
	return distributionConfig
}
func NewClaimConfig() *ClaimConfig {
	claimConfig = &ClaimConfig{
		GlobalConfig: GlobalConfig{
			Config: viper.GetString("config"),
			Debug:  viper.GetBool("debug"),
		},
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		RewardsCoordinatorAddress: viper.GetString("rewards_coordinator_address"),
		Output:                    viper.GetString("output"),
		EarnerAddress:             viper.GetString("earner_address"),
		Tokens:                    viper.GetStringSlice("tokens"),
		ProofStoreBaseUrl:         viper.GetString("proof_store_base_url"),
		ClaimTimestamp:            viper.GetString("claim_timestamp"),
	}
	return claimConfig
}

func getEnvNetwork(environment Environment, network string) (string, error) {
	envString, err := StringEnvironmentFromEnum(environment)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s_%s", envString, network), nil
}

// GetEnvNetwork returns a string concatenation of "{environment}_{network}"
func (c *UpdaterConfig) GetEnvNetwork() (string, error) {
	return getEnvNetwork(c.Environment, c.Network)
}

// GetEnvNetwork returns a string concatenation of "{environment}_{network}"
func (d *DistributionConfig) GetEnvNetwork() (string, error) {
	return getEnvNetwork(d.Environment, d.Network)
}

// GetEnvNetwork returns a string concatenation of "{environment}_{network}"
func (d *ClaimConfig) GetEnvNetwork() (string, error) {
	return getEnvNetwork(d.Environment, d.Network)
}

func KebabToSnakeCase(str string) string {
	return strings.ReplaceAll(str, "-", "_")
}
