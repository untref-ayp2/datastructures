package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDictionary(t *testing.T) {
	dict := NewDictionary[int, int]()

	assert.NotNil(t, dict)
	assert.Equal(t, 0, dict.Size())
}

func TestContains(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)

	assert.True(t, dict.Contains("Leo"))
	assert.True(t, dict.Contains("Lucas"))
	assert.False(t, dict.Contains("Fede"))
}

func TestPut(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)

	assert.Equal(t, 2, dict.Size())
}

func TestPutReplace(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Leo", 38)

	assert.Equal(t, 1, dict.Size())
	assert.Equal(t, 38, dict.Get("Leo"))
}

func TestGet(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Lucas", 35)

	assert.Equal(t, 35, dict.Get("Lucas"))
	assert.Equal(t, 0, dict.Get("Fede"))
}

func TestRemove(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)

	assert.Equal(t, 2, dict.Size())

	dict.Remove("Leo")

	assert.Equal(t, 1, dict.Size())
	assert.True(t, dict.Contains("Lucas"))
	assert.False(t, dict.Contains("Leo"))
}

func TestRemoveNotExists(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)

	assert.Equal(t, 2, dict.Size())

	dict.Remove("Fede")

	assert.Equal(t, 2, dict.Size())
}

func TestSize(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)

	assert.Equal(t, 2, dict.Size())
}

func TestValues(t *testing.T) {
	dic := NewDictionary[int, int]()
	dic.Put(1, 2)
	dic.Put(3, 4)
	dic.Put(5, 6)

	assert.ElementsMatch(t, []int{6, 4, 2}, dic.Values())
}

func TestKeys(t *testing.T) {
	dict := NewDictionary[int, int]()
	dict.Put(1, 2)
	dict.Put(3, 4)
	dict.Put(5, 6)

	assert.ElementsMatch(t, []int{1, 5, 3}, dict.Keys())
}

func TestString(t *testing.T) {
	dict := NewDictionary[int, int]()
	dict.Put(1, 2)
	dict.Put(3, 4)
	dict.Put(5, 6)

	assert.Equal(t, "Dictionary: {\n  1: 2\n  3: 4\n  5: 6\n}", dict.String())
}

func TestStringOnEmptyDictionary(t *testing.T) {
	dict := NewDictionary[int, int]()

	assert.Equal(t, "Dictionary: {}", dict.String())
}
