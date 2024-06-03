package sort

import (
	"testing"

	"github.com/untref-ayp2/data-structures/utils"

	"github.com/stretchr/testify/assert"
)

func TestHeapSortArrayVacio(t *testing.T) {
	arr := make([]int, 0)
	HeapSort[int](arr, utils.Compare[int])
	assert.Equal(t, 0, len(arr))
}
func TestHeapSortArrayConUnSoloElemento(t *testing.T) {
	arr := []int{3}
	HeapSort[int](arr, utils.Compare[int])
	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 3, arr[0])
}
func TestHeapSortDosElementosIguales(t *testing.T) {
	arr := []int{3, 7, 3, -1, -5, 9}
	HeapSort[int](arr, utils.Compare[int])
	assert.Equal(t, []int{-5, -1, 3, 3, 7, 9}, arr)
}
func TestHeapSortArrayMenorAMayor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	HeapSort[int](arr, utils.Compare[int])
	assert.Equal(t, []int{-5, -1, 3, 7, 9}, arr)
}

func TestHeapSortArrayMayorAMenor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	comp := func(a int, b int) int {
		return utils.Compare[int](b, a)
	}
	HeapSort[int](arr, comp)
	assert.Equal(t, []int{9, 7, 3, -1, -5}, arr)
}
