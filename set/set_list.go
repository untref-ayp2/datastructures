package set

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/list"
	"github.com/untref-ayp2/data-structures/types"
)

// SetList implementa un conjunto sobre una lista enlazada simple.
type SetList[T types.Ordered] struct {
	elements *list.LinkedList[T]
}

func NewSetList[T types.Ordered](elements ...T) *SetList[T] {
	set := &SetList[T]{list.NewLinkedList[T]()}
	set.Add(elements...)
	return set
}

func (s *SetList[T]) Contains(element T) bool {
	return s.elements.Find(element) != nil
}

func (s *SetList[T]) Add(elements ...T) {
	for _, element := range elements {
		if !s.Contains(element) {
			s.elements.Append(element)
		}
	}
}

func (s *SetList[T]) Remove(element T) {
	s.elements.Remove(element)
}

func (s *SetList[T]) Size() int {
	return s.elements.Size()
}

func (s *SetList[T]) Values() []T {
	var values []T
	for node := s.elements.Head(); node != nil; node = node.Next() {
		values = append(values, node.Data())
	}
	return values
}

func (s *SetList[T]) String() string {
	result := "Conjunto: {"
	for _, v := range s.Values() {
		result += " " + fmt.Sprintf("%v", v)
	}
	result += " }"
	return result
}
