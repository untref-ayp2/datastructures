package heap

import (
	"errors"

	"github.com/untref-ayp2/data-structures/types"
)

type Heap[T types.Ordered] struct {
	elements  []T
	isMinHeap bool
}

func NewMinHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{isMinHeap: true}
}

func NewMaxHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{isMinHeap: false}
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
		if !m.compare(m.elements[i], m.elements[parent]) {
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

func (m *Heap[T]) compareOrEqual(a T, b T) bool {
	if m.isMinHeap {
		return a <= b
	}

	return a >= b
}

func (m *Heap[T]) compare(a T, b T) bool {
	if m.isMinHeap {
		return a < b
	}

	return a > b
}
