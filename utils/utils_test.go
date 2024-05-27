package utils

import (
	"math"
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

func TestCompare(t *testing.T) {
	assert.Equal(t, Compare(1, 2), -1)
	assert.Equal(t, Compare(2, 1), 1)
	assert.Equal(t, Compare(1, 1), 0)
}

func TestCompareNaN(t *testing.T) {
	assert.Equal(t, Compare(math.NaN(), math.NaN()), 0)
	assert.Equal(t, Compare(math.NaN(), 1), -1)
	assert.Equal(t, Compare(1, math.NaN()), 1)
}
