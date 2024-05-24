package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, Min(1, 2), 1)
	assert.Equal(t, Min(2, 1), 1)
	assert.Equal(t, Min(1, 1), 1)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Max(1, 2), 2)
	assert.Equal(t, Max(2, 1), 2)
	assert.Equal(t, Max(1, 1), 1)
}
