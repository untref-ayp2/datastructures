package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New[int]()
	assert.NotNil(t, s)
}

func TestPush(t *testing.T) {
	s := New[int]()

	s.Push(1)
	assert.False(t, s.IsEmpty())
}

func TestTop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	v, err := s.Top()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
}

func TestTopWhenEmpty(t *testing.T) {
	s := New[int]()

	_, err := s.Top()
	assert.EqualError(t, err, "pila vacía")
}

func TestPop(t *testing.T) {
	s := New[int]()

	s.Push(1)

	v, err := s.Pop()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)
}

func TestPopWhenEmpty(t *testing.T) {
	s := New[int]()

	_, err := s.Pop()
	assert.EqualError(t, err, "pila vacía")
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.False(t, s.IsEmpty())
}
