package stack

import "errors"

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack struct {
	data []any
}

// New crea una nueva pila vacía.
func New() *Stack {
	return &Stack{}
}

// Push agrega un elemento a la pila.
func (s *Stack) Push(x any) {
	s.data = append(s.data, x)
}

// Pop remueve y retorna el elemento en el tope de la pila.
func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("pila vacía")
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, nil
}

// Top retorna el elemento en el tope de la pila.
func (s *Stack) Top() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("pila vacía")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty retorna true si la pila está vacía.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

