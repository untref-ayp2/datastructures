package sort

import (
	"testing"

	"github.com/untref-ayp2/data-structures/utils"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortArrayVacio(t *testing.T) {
	arr := make([]int, 0)
	arr = MergeSort[int](arr, utils.Compare[int])
	assert.Equal(t, 0, len(arr))
}
func TestMergekSortArrayConUnSoloElemento(t *testing.T) {
	arr := []int{3}
	arr = MergeSort[int](arr, utils.Compare[int])
	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 3, arr[0])
}
func TestMergeSortDosElementosIguales(t *testing.T) {
	arr := []int{3, 7, 3, -1, -5, 9}
	arr = MergeSort[int](arr, utils.Compare[int])
	assert.Equal(t, []int{-5, -1, 3, 3, 7, 9}, arr)
}
func TestMergeSortArrayMenorAMayor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	arr = MergeSort[int](arr, utils.Compare[int])
	assert.Equal(t, []int{-5, -1, 3, 7, 9}, arr)
}

func TestMergeSortArrayMayorAMenor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	comp := func(a int, b int) int {
		return utils.Compare[int](b, a)
	}
	arr = MergeSort[int](arr, comp)
	assert.Equal(t, []int{9, 7, 3, -1, -5}, arr)
}
