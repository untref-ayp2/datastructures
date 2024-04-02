package lists

import "cmp"

// Listas de elementos comparables (que implementan la interfaz cmp.Ordered).
// Las listas son estructuras de datos que permiten almacenar y manipular
// colecciones de elementos de manera secuencial. Las operaciones básicas
// que se pueden realizar sobre una lista son: insertar un elemento en una
// posición determinada, eliminar un elemento de una posición determinada,
// obtener un elemento de una posición determinada, reemplazar un elemento
// de una posición determinada, buscar un elemento en la lista y obtener la
// cantidad de elementos en la lista.
type List[T cmp.Ordered] interface{
	// Size retorna la cantidad de elementos en la lista.
	Size() int

	// IsEmpty retorna true si la lista está vacía o false en caso contrario.
	IsEmpty() bool

	// InsertAt agrega un elemento en la posición index de la lista.
	// Si el índice está fuera de rango devuelve un mensaje de error.
	// Si index = 0 inserta al inicio, si index = size inserta al final.
	InsertAt(index int, data T) error

	// RemoveAt elimina el elemento en la posición index de la lista.
	// Si el índice está fuera de rango devuelve un mensaje de error.
	// Si index = 0 elimina al inicio, si index = size-1 elimina el último nodo.
	RemoveAt(index int) (T, error)

	// Get retorna el elemento en la posición index de la lista.
	// Si el índice está fuera de rango devuelve un mensaje de error.
	Get(index int) (T, error)

	// Set reemplaza el elemento en la posición index de la lista por data.
	// Si el índice está fuera de rango devuelve un mensaje de error.
	Set(index int, data T) error

	// IndexOf devuelve la posición de la primera aparción de elemento en la lista.
	// Si el elemento no está en la lista devuelve -1.
	IndexOf(data T) int

	// Clear elimina todos los elementos de la lista.
	Clear()

	// Interate devuelve un canal que permite recorrer la lista.
	Iterate() <-chan T
}