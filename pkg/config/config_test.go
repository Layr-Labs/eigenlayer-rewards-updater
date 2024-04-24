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

func TestGetEnvNetwork_ValidNetworkProd(t *testing.T) {
	envNetwork, err := StringEnvironmentFromEnum(Environment_PROD)
	assert.Nil(t, err)
	assert.Equal(t, "prod", envNetwork)
}

func TestGetEnvNetwork_ValidNetworkPreProd(t *testing.T) {
	envNetwork, err := StringEnvironmentFromEnum(Environment_PRE_PROD)
	assert.Nil(t, err)
	assert.Equal(t, "preprod", envNetwork)
}

func TestGetEnvNetwork_ValidNetworkPreDev(t *testing.T) {
	envNetwork, err := StringEnvironmentFromEnum(Environment_DEV)
	assert.Nil(t, err)
	assert.Equal(t, "dev", envNetwork)
}

func TestGetEnvNetwork_InvalidNetwork(t *testing.T) {
	_, err := StringEnvironmentFromEnum(4)
	assert.NotNil(t, err)
}
