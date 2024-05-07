package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()

	assert.NotNil(t, q)
	assert.True(t, q.IsEmpty())
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)

	v, _ := q.Dequeue()
	assert.Equal(t, 1, v)

	v, _ = q.Dequeue()
	assert.Equal(t, 2, v)

	assert.True(t, q.IsEmpty())
}

func TestQueueDequeueOnEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	_, err := q.Dequeue()
	assert.EqualError(t, err, "cola vacía")
}

func TestQueueFrontOnEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	_, err := q.Front()
	assert.EqualError(t, err, "cola vacía")
}

func TestQueueFront(t *testing.T) {
	q := NewQueue[int]()

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
