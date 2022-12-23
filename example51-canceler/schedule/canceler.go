package schedule

import (
	"context"
	"sync"
	"time"
)

type canceler struct {
	sync.Mutex

	subsciber map[chan struct{}]string
	cancelled map[string]time.Time
}

// Cancel event from api or web
func (c *canceler) Cancel(ctx context.Context, id string) error {
	c.Lock()
	defer c.Unlock()
	c.cancelled[id] = time.Now().Add(5 * time.Minute)
	for subsciber, target := range c.subsciber {
		if id == target {
			close(subsciber)
			return nil
		}
	}
	c.clear()
	return nil
}

// Cancelled connection from worker
func (c *canceler) Cancelled(ctx context.Context, id string) (bool, error) {
	subsciber := make(chan struct{})
	c.Lock()
	c.subsciber[subsciber] = id
	c.Unlock()

	defer func() {
		c.Lock()
		delete(c.subsciber, subsciber)
		c.Unlock()
	}()

	c.Lock()
	_, ok := c.cancelled[id]
	c.Unlock()
	if ok {
		return true, nil
	}

	select {
	case <-ctx.Done():
		return false, nil
	case <-subsciber:
		return true, nil
	}
}

func (c *canceler) clear() {
	now := time.Now()
	for k, trigger := range c.cancelled {
		if now.After(trigger) {
			delete(c.cancelled, k)
		}
	}
}

func newCanceler() *canceler {
	return &canceler{
		subsciber: make(map[chan struct{}]string),
		cancelled: make(map[string]time.Time),
	}
}
