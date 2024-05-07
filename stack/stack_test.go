package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.NotNil(t, s)
}

func TestStackPush(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	assert.False(t, s.IsEmpty())
}

func TestStackTop(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)
	v, err := s.Top()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
}

func TestStackTopWhenEmpty(t *testing.T) {
	s := NewStack[int]()

	_, err := s.Top()
	assert.EqualError(t, err, "pila vacía")
}

func TestStackPop(t *testing.T) {
	s := NewStack[int]()

	s.Push(1)

	v, err := s.Pop()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
}

func TestStackPopWhenEmpty(t *testing.T) {
	s := NewStack[int]()

	_, err := s.Pop()
	assert.EqualError(t, err, "pila vacía")
}

func TestStackIsEmpty(t *testing.T) {
	s := NewStack[int]()
	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.False(t, s.IsEmpty())
}
