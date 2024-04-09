package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.False(t, q.IsEmpty())

	v, _ := q.Dequeue()
	assert.Equal(t, 1, v)

	v, _ = q.Dequeue()
	assert.Equal(t, 2, v)

	v, _ = q.Dequeue()
	assert.Equal(t, 3, v)

	_, err := q.Dequeue()
	assert.Error(t, err, "cola vacía")
}

func TestEmptyQueue(t *testing.T) {
	q := New[int]()

	assert.True(t, q.IsEmpty())

	_, err := q.Dequeue()
	assert.Error(t, err, "cola vacía")

	_, err = q.Front()
	assert.Error(t, err, "cola vacía")
}

func TestFront(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, _ := q.Front()
	assert.Equal(t, 1, v)

	q.Dequeue()

	v, _ = q.Front()
	assert.Equal(t, 2, v)

	q.Dequeue()

	v, _ = q.Front()
	assert.Equal(t, 3, v)
}
