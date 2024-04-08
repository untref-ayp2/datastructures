package set

import (
	"fmt"

	sl "github.com/untref-ayp2/data-structures/lists/single_linked_list"
	"github.com/untref-ayp2/data-structures/types"
)

type Set[T types.Ordered] struct {
	elements sl.SingleLinkedList[T]
}

func (s *Set[T]) Contains(element T) bool {
	return s.elements.Find(element) != nil
}

func (s *Set[T]) Add(element T) {
	if !s.Contains(element) {
		s.elements.Append(element)
	}
}

func (s *Set[T]) Remove(element T) {
	s.elements.Remove(element)
}

func (s *Set[T]) Size() int {
	return s.elements.Size()
}

func (s *Set[T]) Values() []T {
	var values []T
	i := 0
	for node := s.elements.Head(); node != nil; node = node.Next() {
		values = append(values, node.Data())
		i++
	}
	return values
}

func NewSet[T types.Ordered](elements ...T) *Set[T] {
	set := &Set[T]{*sl.NewList[T]()}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

func (s *Set[T]) String() string {
	result := "Conjunto: {"
	for _, v := range s.Values() {
		result += " " + fmt.Sprintf("%v", v)
	}
	result += " }"
	return result
}

func Union[T types.Ordered](s1, s2 *Set[T]) *Set[T] {
	set := NewSet[T]()
	for _, v := range s1.Values() {
		set.Add(v)
	}
	for _, v := range s2.Values() {
		set.Add(v)
	}
	return set
}

func Intersection[T types.Ordered](s1, s2 *Set[T]) *Set[T] {
	set := NewSet[T]()
	for _, v := range s1.Values() {
		if s2.Contains(v) {
			set.Add(v)
		}
	}
	return set
}

func Difference[T types.Ordered](s1, s2 *Set[T]) *Set[T] {
	set := NewSet[T]()
	for _, v := range s1.Values() {
		if !s2.Contains(v) {
			set.Add(v)
		}
	}
	return set
}

func Subset[T types.Ordered](s1, s2 *Set[T]) bool {
	for _, v := range s2.Values() {
		if !s1.Contains(v) {
			return false
		}
	}
	return true
}

func Equal[T types.Ordered](s1, s2 *Set[T]) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}
