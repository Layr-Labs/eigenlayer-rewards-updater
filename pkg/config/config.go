package config

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Config             string `mapstructure:"config"`
	Debug              bool   `mapstructure:"debug"`
	DDStatsdUrl        string `mapstructure:"dd_statsd_url"`
	EnableStatsd       bool   `mapstructure:"enable_statsd"`
	PushgatewayUrl     string `mapstructure:"pushgateway_url"`
	PushgatewayEnabled bool   `mapstructure:"pushgateway_enabled"`
	EnableTracing      bool   `mapstructure:"enable_tracing"`
}

type Environment int

var (
	Environment_PRE_PROD Environment = 0
	Environment_TESTNET  Environment = 1
	Environment_MAINNET  Environment = 2
)

type UpdaterConfig struct {
	GlobalConfig
	Environment               Environment `mapstructure:"environment"`
	Network                   string      `mapstructure:"network"`
	RPCUrl                    string      `mapstructure:"rpc_url"`
	PrivateKey                string      `mapstructure:"private_key"`
	RewardsCoordinatorAddress string      `mapstructure:"rewards_coordinator_address"`
	SidecarRpcUrl             string      `mapstructure:"sidecar_rpc_url"`
	SidecarInsecureRpc        bool        `mapstructure:"sidecar_insecure_rpc"`
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
	SubmitClaim               bool        `mapstructure:"submit_claim"`
	SidecarRpcUrl             string      `mapstructure:"sidecar_rpc_url"`
	SidecarInsecureRpc        bool        `mapstructure:"sidecar_insecure_rpc"`
}

type ValidateConfig struct {
	GlobalConfig
	Environment               Environment `mapstructure:"environment"`
	Network                   string      `mapstructure:"network"`
	RPCUrl                    string      `mapstructure:"rpc_url"`
	PrivateKey                string      `mapstructure:"private_key"`
	RewardsCoordinatorAddress string      `mapstructure:"rewards_coordinator_address"`
	ProofStoreBaseUrl         string      `mapstructure:"proof_store_base_url"`
}

type DisableRootConfig struct {
	GlobalConfig
	Environment               Environment `mapstructure:"environment"`
	Network                   string      `mapstructure:"network"`
	RPCUrl                    string      `mapstructure:"rpc_url"`
	PrivateKey                string      `mapstructure:"private_key"`
	RewardsCoordinatorAddress string      `mapstructure:"rewards_coordinator_address"`
	ProofStoreBaseUrl         string      `mapstructure:"proof_store_base_url"`
	RootIndex                 uint32      `mapstructure:"root_index"`
}

var updaterConfig *UpdaterConfig
var distributionConfig *DistributionConfig
var claimConfig *ClaimConfig
var validateConfig *ValidateConfig
var disableRootConfig *DisableRootConfig

// parseEnvironment normalizes environment names to an enum value
func parseEnvironment(env string) Environment {
	switch env {
	case "preprod":
		return Environment_PRE_PROD
	case "testnet":
		return Environment_TESTNET
	case "mainnet":
		return Environment_MAINNET
	default:
		return Environment_PRE_PROD
	}
}

// StringEnvironmentFromEnum gets a string environment value from the enum
func StringEnvironmentFromEnum(env Environment) (string, error) {
	switch env {
	case Environment_PRE_PROD:
		return "preprod", nil
	case Environment_TESTNET:
		return "testnet", nil
	case Environment_MAINNET:
		return "mainnet", nil
	}
	return "", errors.New(fmt.Sprintf("String env not found for '%d'", env))
}

func GetGlobalConfig() GlobalConfig {
	return GlobalConfig{
		Config:             viper.GetString("config"),
		Debug:              viper.GetBool("debug"),
		DDStatsdUrl:        viper.GetString("dd_statsd_url"),
		EnableTracing:      viper.GetBool("enable_tracing"),
		PushgatewayEnabled: viper.GetBool("pushgateway_enabled"),
		PushgatewayUrl:     viper.GetString("pushgateway_url"),
	}
}

// NewUpdaterConfig reads config values from viper and returns
// them in a struct
func NewUpdaterConfig() *UpdaterConfig {
	updaterConfig = &UpdaterConfig{
		GlobalConfig:              GetGlobalConfig(),
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		RewardsCoordinatorAddress: viper.GetString("rewards_coordinator_address"),
		SidecarRpcUrl:             viper.GetString("sidecar_rpc_url"),
		SidecarInsecureRpc:        viper.GetBool("sidecar_insecure_rpc"),
	}
	return updaterConfig
}

func NewDistributionConfig() *DistributionConfig {
	distributionConfig = &DistributionConfig{
		GlobalConfig:              GetGlobalConfig(),
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
		GlobalConfig:              GetGlobalConfig(),
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
		SubmitClaim:               viper.GetBool("submit_claim"),
		SidecarRpcUrl:             viper.GetString("sidecar_rpc_url"),
		SidecarInsecureRpc:        viper.GetBool("sidecar_insecure_rpc"),
	}
	return claimConfig
}
func NewValidateConfig() *ValidateConfig {
	validateConfig = &ValidateConfig{
		GlobalConfig:              GetGlobalConfig(),
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		RewardsCoordinatorAddress: viper.GetString("rewards_coordinator_address"),
		ProofStoreBaseUrl:         viper.GetString("proof_store_base_url"),
	}
	return validateConfig
}

func NewDisableRootConfig() *DisableRootConfig {
	disableRootConfig = &DisableRootConfig{
		GlobalConfig:              GetGlobalConfig(),
		Environment:               parseEnvironment(viper.GetString("environment")),
		Network:                   viper.GetString("network"),
		RPCUrl:                    viper.GetString("rpc_url"),
		PrivateKey:                viper.GetString("private_key"),
		RewardsCoordinatorAddress: viper.GetString("rewards_coordinator_address"),
		ProofStoreBaseUrl:         viper.GetString("proof_store_base_url"),
		RootIndex:                 viper.GetUint32("root_index"),
	}
	return disableRootConfig
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
