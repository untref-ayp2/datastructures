package heap

import (
	"errors"

	"github.com/untref-ayp2/data-structures/types"
)

type MaxHeap[T types.Ordered] struct {
	elements []T
}

func NewMaxHeap[T types.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{}
}

func (m *MaxHeap[T]) Len() int {
	return len(m.elements)
}

func (m *MaxHeap[T]) Insert(element T) {
	m.elements = append(m.elements, element)
	m.upHeap(len(m.elements) - 1)
}

func (m *MaxHeap[T]) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if m.elements[i] <= m.elements[parent] {
			break
		}
		m.elements[i], m.elements[parent] = m.elements[parent], m.elements[i]
		i = parent
	}
}

func (m *MaxHeap[T]) RemoveMax() (T, error) {
	var element T
	if m.Len() == 0 {
		return element, errors.New("heap vacío")
	}
	element = m.elements[0]
	m.elements[0] = m.elements[m.Len()-1]
	m.elements = m.elements[:m.Len()-1]
	m.downHeap(0)

	return element, nil
}

func (m *MaxHeap[T]) downHeap(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		if left >= m.Len() {
			break
		}
		max := left
		if right < m.Len() && m.elements[right] > m.elements[left] {
			max = right
		}
		if m.elements[i] >= m.elements[max] {
			break
		}
		m.elements[i], m.elements[max] = m.elements[max], m.elements[i]
		i = max
	}
}
