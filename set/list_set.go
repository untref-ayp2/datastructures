package set

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/list"
)

// ListSet implementa un conjunto sobre una lista enlazada simple.
type ListSet[T comparable] struct {
	elements *list.LinkedList[T]
}

func NewListSet[T comparable](elements ...T) *ListSet[T] {
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
	str := "Set: {"
	for i, v := range s.Values() {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", v)

	}
	str += "}"
	return str
}
