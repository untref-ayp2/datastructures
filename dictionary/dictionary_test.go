package dictionary

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDictionary(t *testing.T) {
	dic := NewDictionary[int, int]()

	assert.NotNil(t, dic)
	assert.Equal(t, 0, dic.Size())
}
func TestContains(t *testing.T) {
	dic := NewDictionary[string, int]()
	dic.Put("Leo", 55)
	dic.Put("lucas", 38)

	assert.True(t, dic.Contains("Leo"))
	assert.True(t, dic.Contains("lucas"))
	assert.False(t, dic.Contains("Fede"))
}

func TestPut(t *testing.T) {
	dic := NewDictionary[string, int]()
	dic.Put("Leo", 55)
	dic.Put("Lucas", 38)

	assert.Equal(t, 2, dic.Size())
}

func TestPutEqualsElements(t *testing.T) {
	dic := NewDictionary[string, int]()
	dic.Put("Leo", 55)
	dic.Put("Leo", 38)

	assert.Equal(t, 1, dic.Size())
	assert.Equal(t, 38, dic.Get("Leo"))
}

func TestGet(t *testing.T) {
	dic := NewDictionary[string, int]()
	dic.Put("Lucas", 35)

	assert.Equal(t, 35, dic.Get("Lucas"))
	assert.Equal(t, 0, dic.Get("Fede"))
}

func TestRemove(t *testing.T) {
	dic := NewDictionary[string, int]()
	dic.Put("Leo", 55)
	dic.Put("Lucas", 38)

	assert.Equal(t, 2, dic.Size())

	dic.Remove("Leo")

	assert.Equal(t, 1, dic.Size())
	assert.True(t, dic.Contains("Lucas"))
	assert.False(t, dic.Contains("Leo"))
}

func TestGetValues(t *testing.T) {
	dic := NewDictionary[int, int]()
	dic.Put(1, 2)
	dic.Put(3, 4)
	dic.Put(5, 6)
	esperado := []int{2, 4, 6}
	obtenido := dic.GetValues()
	sort.Ints(obtenido)

	assert.Equal(t, esperado, obtenido)
}

func TestGetKeys(t *testing.T) {
	dic := NewDictionary[int, int]()
	dic.Put(1, 2)
	dic.Put(3, 4)
	dic.Put(5, 6)
	esperado := []int{1, 3, 5}
	obtenido := dic.GetKeys()
	sort.Ints(dic.GetKeys())

	assert.Equal(t, esperado, obtenido)
}
