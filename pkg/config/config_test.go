package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKebabToSnakeCase(t *testing.T) {
	kebabString := "some-kebab-string"

	snakeString := KebabToSnakeCase(kebabString)

	assert.Equal(t, "some_kebab_string", snakeString)
}

func TestGetEnvNetwork_ValidNetworkPreProd(t *testing.T) {
	envNetwork, err := StringEnvironmentFromEnum(Environment_PRE_PROD)
	assert.Nil(t, err)
	assert.Equal(t, "preprod", envNetwork)
}

func TestGetEnvNetwork_ValidNetworkMainnet(t *testing.T) {
	envNetwork, err := StringEnvironmentFromEnum(Environment_MAINNET)
	assert.Nil(t, err)
	assert.Equal(t, "mainnet", envNetwork)
}

func TestGetEnvNetwork_InvalidNetwork(t *testing.T) {
	_, err := StringEnvironmentFromEnum(5)
	assert.NotNil(t, err)
}
