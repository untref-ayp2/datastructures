package stack

import "errors"

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack[T any] struct {
	data []T
}

// New crea una nueva pila vacía.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push agrega un elemento a la pila.
func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

// Pop remueve y retorna el elemento en el tope de la pila.
func (s *Stack[T]) Pop() (T, error) {
	var x T
	if s.IsEmpty() {
		return x, errors.New("pila vacía")
	}
	x = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, nil
}

// Top retorna el elemento en el tope de la pila.
func (s *Stack[T]) Top() (T, error) {
	var x T
	if s.IsEmpty() {
		return x, errors.New("pila vacía")
	}
	x = s.data[len(s.data)-1]
	return x, nil
}

// IsEmpty retorna true si la pila está vacía.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
