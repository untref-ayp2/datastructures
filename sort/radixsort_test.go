package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadixSorttVacio(t *testing.T) {
	arr := []string{}
	RadixSort(arr)
	assert.Equal(t, []string{}, arr)
}

func TestRadixSorttUnicoElemento(t *testing.T) {
	arr := []string{"123"}
	RadixSort(arr)
	assert.Equal(t, []string{"123"}, arr)
}

func TestRadixSorttInt(t *testing.T) {
	arr := []string{"123", "12"}
	RadixSort(arr)
	assert.Equal(t, []string{"12", "123"}, arr)
}
func TestRadixSorttString(t *testing.T) {
	arr := []string{"ab", "aAa", "a", "ac", "aaa", "aca"}
	RadixSort(arr)
	assert.Equal(t, []string{"a", "aAa", "aaa", "ab", "ac", "aca"}, arr)
}

func TestRadixSortt(t *testing.T) {
	arr := []string{"123abc", "120abc", "120aac", "120", "12a"}
	RadixSort(arr)
	assert.Equal(t, []string{"120", "120aac", "120abc", "123abc", "12a"}, arr)
}
