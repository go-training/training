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
	for subsciber, target := range c.subsciber {
		if id == target {
			close(subsciber)
			delete(c.subsciber, subsciber)
		}
	}
	return nil
}

func newCanceler() *canceler {
	return &canceler{
		subsciber: make(map[chan struct{}]string),
	}
}
