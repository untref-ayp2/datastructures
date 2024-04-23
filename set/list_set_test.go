package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := NewListSet[int]()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestAdd(t *testing.T) {
	set := NewListSet[int]()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestAddMultiple(t *testing.T) {
	set := NewListSet[int]()

	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
}

func TestAddExistenteNoRepite(t *testing.T) {
	set := NewListSet[int]()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestContains(t *testing.T) {
	set := NewListSet[int]()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestRemove(t *testing.T) {
	set := NewListSet[int]()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestRemoveNonExistent(t *testing.T) {
	set := NewListSet[int]()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestSize(t *testing.T) {
	set := NewListSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestValuesOnAnEmptySet(t *testing.T) {
	set := NewListSet[int]()
	values := set.Values()

	assert.Equal(t, 0, len(values))
}

func TestValuesOnANonEmptySet(t *testing.T) {
	set := NewListSet(1, 2)
	values := set.Values()

	assert.Equal(t, 2, len(values))
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestStringEnSetVacio(t *testing.T) {
	set := NewListSet[int]()
	assert.Equal(t, "Set: {}", set.String())
}

func TestStringEnSetNoVacio(t *testing.T) {
	set := NewListSet(1, 2)
	assert.Regexp(t, `Set: \{[12], [12]\}`, set.String())
}
