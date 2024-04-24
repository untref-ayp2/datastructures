package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListSet(t *testing.T) {
	set := NewListSet[int]()

	assert.NotNil(t, set)
	assert.Equal(t, 0, set.Size())
}

func TestListSetAdd(t *testing.T) {
	set := NewListSet[int]()

	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestListSetAddMultiple(t *testing.T) {
	set := NewListSet[int]()

	set.Add(1, 2, 3)
	assert.Equal(t, 3, set.Size())
}

func TestListSetAddExistenteNoRepite(t *testing.T) {
	set := NewListSet[int]()

	set.Add(1)
	set.Add(1)
	assert.Equal(t, 1, set.Size())
}

func TestListSetContains(t *testing.T) {
	set := NewListSet[int]()
	set.Add(1)

	assert.True(t, set.Contains(1))
	assert.False(t, set.Contains(2))
}

func TestListSetRemove(t *testing.T) {
	set := NewListSet[int]()
	set.Add(1)
	set.Add(2)
	assert.True(t, set.Contains(1))
	assert.Equal(t, 2, set.Size())

	set.Remove(1)
	assert.False(t, set.Contains(1))
	assert.Equal(t, 1, set.Size())
}

func TestListSetRemoveNonExistent(t *testing.T) {
	set := NewListSet[int]()
	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Remove(2)
	assert.Equal(t, 1, set.Size())
}

func TestListSetSize(t *testing.T) {
	set := NewListSet[int]()
	assert.Equal(t, 0, set.Size())

	set.Add(1)
	assert.Equal(t, 1, set.Size())

	set.Add(2)
	assert.Equal(t, 2, set.Size())
}

func TestListSetValuesOnAnEmptySet(t *testing.T) {
	set := NewListSet[int]()
	values := set.Values()

	assert.Empty(t, values)
}

func TestListSetValuesOnANonEmptySet(t *testing.T) {
	set := NewListSet(1, 2)
	values := set.Values()

	assert.Len(t, values, 2)
	assert.ElementsMatch(t, []int{1, 2}, values)
}

func TestListSetStringEnSetVacio(t *testing.T) {
	set := NewListSet[int]()
	assert.Equal(t, "Set: {}", set.String())
}

func TestListSetStringEnSetNoVacio(t *testing.T) {
	set := NewListSet(1, 2)
	possibleRepresentations := []string{
		"Set: {1, 2}",
		"Set: {2, 1}",
	}

	assert.Contains(t, possibleRepresentations, set.String())
}
