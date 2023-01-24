package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	size := 100
	q := NewCircularBuffer(size)

	for i := 0; i < (size - 1); i++ {
		assert.NoError(t, q.Enqueue(i+1))
	}
	// can't insert new data.
	assert.Error(t, q.Enqueue(4))

	for i := 0; i < (size - 1); i++ {
		v, err := q.Dequeue()
		assert.Equal(t, i+1, v.(int))
		assert.NoError(t, err)
	}
}
