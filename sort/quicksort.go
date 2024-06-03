package sort

// QuickSort ordena un array de elementos de cualquier tipo utilizando el algoritmo de quicksort.
// Recibe un array de elementos, una función de comparación y un entero que indica el orden de ordenamiento.
// Para ordenar de mayor a menor la función de comparación debe devolver:
//
//	-1 si el primer elemento es menor que el segundo,
//	 0 si ambos elementos son iguales,
//	+1 si el primer elemento es mayor que el segundo.
func QuickSort[T any](arr []T, compare func(a T, b T) int) {
	quickSort(arr, compare, 0, len(arr)-1)
}

func quickSort[T any](arr []T, compare func(a T, b T) int, low, high int) {
	if low < high {
		pi := partition(arr, compare, low, high)

		quickSort(arr, compare, low, pi-1)
		quickSort(arr, compare, pi+1, high)
	}
}

func partition[T any](arr []T, compare func(a T, b T) int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if compare(arr[j], pivot) < 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
