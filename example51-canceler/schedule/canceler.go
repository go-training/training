package schedule

import (
	"context"
	"sync"
)

type canceler struct {
	sync.Mutex

	subsciber map[chan struct{}]string
}

// Cancel event from api or web
func (c *canceler) Cancel(ctx context.Context, id string) error {
	c.Lock()
	defer c.Unlock()
	for subsciber, target := range c.subsciber {
		if id == target {
			close(subsciber)
		}
	}
	return nil
}

// Canceled connection from worker
func (c *canceler) Canceled(ctx context.Context, id string) (bool, error) {
	subsciber := make(chan struct{})
	c.Lock()
	c.subsciber[subsciber] = id
	c.Unlock()

	defer func() {
		c.Lock()
		delete(c.subsciber, subsciber)
		c.Unlock()
	}()

	select {
	case <-ctx.Done():
		return false, nil
	case <-subsciber:
		return true, nil
	}
}

func newCanceler() *canceler {
	return &canceler{
		subsciber: make(map[chan struct{}]string),
	}
}
