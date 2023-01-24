package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	size := 100
	q := NewCircularBuffer(size)

	for i := 0; i < size; i++ {
		assert.NoError(t, q.Enqueue(i))
	}
	// can't insert new data.
	assert.Error(t, q.Enqueue(0))
	assert.Equal(t, errFull, q.Enqueue(0))

	for i := 0; i < size; i++ {
		v, err := q.Dequeue()
		assert.Equal(t, i, v.(int))
		assert.NoError(t, err)
	}

	// no task
	_, err := q.Dequeue()
	assert.Error(t, err)
	assert.Equal(t, errNoTask, err)
}

func BenchmarkCircularBufferEnqueueDequeue(b *testing.B) {
	q := NewCircularBuffer(b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Enqueue(i)
		_, _ = q.Dequeue()
	}
}

func BenchmarkCircularBufferEnqueue(b *testing.B) {
	q := NewCircularBuffer(b.N)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = q.Enqueue(i)
	}
}

func BenchmarkCircularBufferDequeue(b *testing.B) {
	q := NewCircularBuffer(b.N)

	for i := 0; i < b.N; i++ {
		_ = q.Enqueue(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = q.Dequeue()
	}
}
