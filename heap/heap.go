package heap

import (
	"cmp"
	"errors"

	"github.com/untref-ayp2/data-structures/types"
)

type Heap[T any] struct {
	// contenedor de datos
	elements  []T
	//Función de comparación. Para un heap de mínimo,
	// devuelve 1 si a < b, 0 si a == b, -1 si a > b
	// Para un heap de máximo, devuelve 1 si a > b, 0 si a == b, -1 si a < b
	comp func(a T, b T) int  
}

func NewMinHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{comp: cmp.Compare[T], elements: make([]T, 0)}
}

func NewMaxHeap[T types.Ordered]() *Heap[T] {
	comp:= func(a T, b T) int {
		return cmp.Compare[T](b, a)
	}
	return &Heap[T]{comp: comp, elements: make([]T, 0)}
}

func NewGenericHeap[T any](comp func(a T, b T) int) *Heap[T] {
	return &Heap[T]{comp: comp, elements: make([]T, 0)}
}

func (m *Heap[T]) Size() int {
	return len(m.elements)
}

func (m *Heap[T]) Insert(element T) {
	m.elements = append(m.elements, element)
	m.upHeap(len(m.elements) - 1)
}

func (m *Heap[T]) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if m.comp(m.elements[i], m.elements[parent]) == 1 {
			break
		}
		m.elements[i], m.elements[parent] = m.elements[parent], m.elements[i]
		i = parent
	}
}

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

func (m *Heap[T]) downHeap(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < m.Size() && m.comp(m.elements[left], m.elements[smallest]) == -1 {
			smallest = left
		}

		if right < m.Size() && m.comp(m.elements[right], m.elements[smallest]) == -1 {
			smallest = right
		}

		if smallest == i {
			break
		}

		m.elements[i], m.elements[smallest] = m.elements[smallest], m.elements[i]
		i = smallest
	}
}
