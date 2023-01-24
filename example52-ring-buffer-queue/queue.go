package queue

import (
	"errors"
	"sync"
)

type T interface{}

var (
	errFull   = errors.New("full")
	errNoTask = errors.New("no task")
)

type CircularBuffer struct {
	sync.Mutex
	taskQueue []T
	capacity  int
	head      int
	tail      int
	full      bool
}

func (s *CircularBuffer) IsEmpty() bool {
	return s.head == s.tail && !s.full
}

func (s *CircularBuffer) IsFull() bool {
	return s.full
}

func (s *CircularBuffer) Enqueue(task T) error {
	if s.IsFull() {
		return errFull
	}

	s.Lock()
	s.taskQueue[s.tail] = task
	s.tail = (s.tail + 1) % s.capacity
	s.full = s.head == s.tail
	s.Unlock()

	return nil
}

func (s *CircularBuffer) Dequeue() (T, error) {
	if s.IsEmpty() {
		return nil, errNoTask
	}

	s.Lock()
	data := s.taskQueue[s.head]
	s.full = false
	s.head = (s.head + 1) % s.capacity
	s.Unlock()

	return data, nil
}

func NewCircularBuffer(size int) *CircularBuffer {
	w := &CircularBuffer{
		taskQueue: make([]T, size),
		capacity:  size,
	}

	return w
}
