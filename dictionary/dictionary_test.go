package dictionary

import (
	"sort"
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
	dict.Put("lucas", 38)

	assert.True(t, dict.Contains("Leo"))
	assert.True(t, dict.Contains("lucas"))
	assert.False(t, dict.Contains("Fede"))
}

func TestPut(t *testing.T) {
	dict := NewDictionary[string, int]()
	dict.Put("Leo", 55)
	dict.Put("Lucas", 38)

	assert.Equal(t, 2, dict.Size())
}

func TestPutEqualsElements(t *testing.T) {
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

func TestValues(t *testing.T) {
	dic := NewDictionary[int, int]()
	dic.Put(1, 2)
	dic.Put(3, 4)
	dic.Put(5, 6)
	esperado := []int{2, 4, 6}
	obtenido := dic.Values()
	sort.Ints(obtenido)

	assert.Equal(t, esperado, obtenido)
}

func TestKeys(t *testing.T) {
	dict := NewDictionary[int, int]()
	dict.Put(1, 2)
	dict.Put(3, 4)
	dict.Put(5, 6)
	esperado := []int{1, 3, 5}
	obtenido := dict.Keys()
	sort.Ints(dict.Keys())

	assert.Equal(t, esperado, obtenido)
}
