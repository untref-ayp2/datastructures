package sort

import (
	"testing"

	"github.com/untref-ayp2/data-structures/utils"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortArrayVacio(t *testing.T) {
	arr := make([]int, 0)
	QuickSort[int](arr, utils.Compare[int])
	assert.Equal(t, 0, len(arr))
}
func TestQuickSortArrayConUnSoloElemento(t *testing.T) {
	arr := []int{3}
	QuickSort[int](arr, utils.Compare[int])
	assert.Equal(t, 1, len(arr))
	assert.Equal(t, 3, arr[0])
}
func TestQuickSortDosElementosIguales(t *testing.T) {
	arr := []int{3, 7, 3, -1, -5, 9}
	QuickSort[int](arr, utils.Compare[int])
	assert.Equal(t, []int{-5, -1, 3, 3, 7, 9}, arr)
}
func TestQuickSortArrayMenorAMayor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	QuickSort[int](arr, utils.Compare[int])
	assert.Equal(t, []int{-5, -1, 3, 7, 9}, arr)
}

func TestQuickSortArrayMayorAMenor(t *testing.T) {
	arr := []int{3, 7, -1, -5, 9}
	comp := func(a int, b int) int {
		return utils.Compare[int](b, a)
	}
	QuickSort[int](arr, comp)
	assert.Equal(t, []int{9, 7, 3, -1, -5}, arr)
}
