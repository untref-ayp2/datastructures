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

func TestRadixSorttFloat(t *testing.T) {
	arr := []string{"123.5", "12", "123", "0.201", "0.200"}
	RadixSort(arr)
	assert.Equal(t, []string{"0.200", "0.201", "12", "123", "123.5"}, arr)
}

func TestRadixSorttñ(t *testing.T) {
	arr := []string{"caño", "año", "ñandu", "niño", "niña"}
	RadixSort(arr)
	assert.Equal(t, []string{"año", "caño", "niña", "niño", "ñandu"}, arr)
}
