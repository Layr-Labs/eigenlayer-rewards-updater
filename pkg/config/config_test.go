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
