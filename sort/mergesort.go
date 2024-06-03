package sort

func MergeSort[T any](arr []T, compare func(T, T) int) []T {

	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle], compare)
	right := MergeSort(arr[middle:], compare)

	return merge(left, right, compare)
}

func merge[T any](left, right []T, compare func(T, T) int) []T {

	size, i, j := len(left)+len(right), 0, 0
	lenRight, lenLeft := len(right)-1, len(left)-1
	array := make([]T, size)

	for k := 0; k < size; k++ {
		if i > lenLeft && j <= lenRight {
			array[k] = right[j]
			j++
		} else if j > lenRight && i <= lenLeft {
			array[k] = left[i]
			i++
		} else if compare(left[i], right[j]) < 0 {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
	}

	return array
}
