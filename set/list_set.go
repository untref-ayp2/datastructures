package set

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/list"
	"github.com/untref-ayp2/data-structures/types"
)

// ListSet implementa un conjunto sobre una lista enlazada simple.
type ListSet[T types.Ordered] struct {
	elements *list.LinkedList[T]
}

func NewListSet[T types.Ordered](elements ...T) *ListSet[T] {
	set := &ListSet[T]{list.NewLinkedList[T]()}
	set.Add(elements...)
	return set
}

func (s *ListSet[T]) Contains(element T) bool {
	return s.elements.Find(element) != nil
}

func (s *ListSet[T]) Add(elements ...T) {
	for _, element := range elements {
		if !s.Contains(element) {
			s.elements.Append(element)
		}
	}
}

func (s *ListSet[T]) Remove(element T) {
	s.elements.Remove(element)
}

func (s *ListSet[T]) Size() int {
	return s.elements.Size()
}

func (s *ListSet[T]) Values() []T {
	var values []T
	for node := s.elements.Head(); node != nil; node = node.Next() {
		values = append(values, node.Data())
	}
	return values
}

func (s *ListSet[T]) String() string {
	result := "Conjunto: {"
	for _, v := range s.Values() {
		result += " " + fmt.Sprintf("%v", v)
	}
	result += " }"
	return result
}
