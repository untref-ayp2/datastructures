// Package heap provee una implementación de un heap binario.
package heap

import (
	"errors"

	"github.com/untref-ayp2/data-structures/types"
)

type Heap[T types.Ordered] struct {
	elements  []T
	isMinHeap bool
}

// NewMinHeap crea un nuevo heap binario de mínimos.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//
// Retorna:
//   - un puntero a un heap binario de mínimos.
func NewMinHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{isMinHeap: true}
}

// NewMaxHeap crea un nuevo heap binario de máximos.
//
// Uso:
//
//	heap := heap.NewMaxHeap[int]()
//
// Retorna:
//   - un puntero a un heap binario de máximos.
func NewMaxHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{isMinHeap: false}
}

// Size retorna la cantidad de elementos en el heap.
//
// Uso:
//
//	size := heap.Size()
//
// Retorna:
//   - la cantidad de elementos en el heap.
func (m *Heap[T]) Size() int {
	return len(m.elements)
}

// Insert agrega un elemento al heap.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//	heap.Insert(5)
//
// Parámetros:
//
//	element: elemento a agregar al heap.
func (m *Heap[T]) Insert(element T) {
	m.elements = append(m.elements, element)
	m.upHeap(len(m.elements) - 1)
}

// upHeap reordena el heap hacia arriba.
//
// Parámetros:
//   - `i` índice del elemento a reordenar.
func (m *Heap[T]) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if !m.compare(m.elements[i], m.elements[parent]) {
			break
		}
		m.elements[i], m.elements[parent] = m.elements[parent], m.elements[i]
		i = parent
	}
}

// Remove elimina y retorna el elemento en la cima del heap.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//	heap.Insert(5)
//	element, _ := heap.Remove()
//
// Retorna:
//   - el elemento en la cima del heap.
func (m *Heap[T]) Remove() (T, error) {
	var element T
	if m.Size() == 0 {
		return element, errors.New("heap vacío")
	}
	element = m.elements[0]
	m.elements[0] = m.elements[m.Size()-1]
	m.elements = m.elements[:m.Size()-1]
	m.downHeap(0)

	return element, nil
}

// downHeap reordena el heap hacia abajo.
//
// Parámetros:
//   - `i` índice del elemento a reordenar.
func (m *Heap[T]) downHeap(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		if left >= m.Size() {
			break
		}
		aux := left
		if right < m.Size() && m.compare(m.elements[right], m.elements[left]) {
			aux = right
		}
		if m.compareOrEqual(m.elements[i], m.elements[aux]) {
			break
		}
		m.elements[i], m.elements[aux] = m.elements[aux], m.elements[i]
		i = aux
	}
}

// compareOrEqual compara o verifica igualdad de dos elementos de acuerdo al tipo de heap.
//
// Parámetros:
//   - `a` primer elemento a comparar.
//   - `b` segundo elemento a comparar.
//
// Retorna:
//   - true si `a` es menor/mayor o igual a `b`, false en caso contrario.
func (m *Heap[T]) compareOrEqual(a T, b T) bool {
	if m.isMinHeap {
		return a <= b
	}

	return a >= b
}

// compare compara dos elementos de acuerdo al tipo de heap.
//
// Parámetros:
//   - `a` primer elemento a comparar.
//   - `b` segundo elemento a comparar.
//
// Retorna:
//   - true si `a` es menor/mayor a `b`, false en caso contrario.
func (m *Heap[T]) compare(a T, b T) bool {
	if m.isMinHeap {
		return a < b
	}

	return a > b
}
