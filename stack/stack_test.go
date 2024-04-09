package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()
	assert.Equal(t, 3, v)
	assert.NoError(t, err)

	v, err = s.Pop()
	assert.Equal(t, 2, v)
	assert.NoError(t, err)

	v, err = s.Pop()
	assert.Equal(t, 1, v)
	assert.NoError(t, err)

	_, err = s.Pop()
	assert.Error(t, err, "pila vacía")
}
