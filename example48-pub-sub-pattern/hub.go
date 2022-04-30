package main

import (
	"context"
	"sync"
)

type hub struct {
	sync.Mutex
	subs map[*subscriber]struct{}
}

func (h *hub) publish(ctx context.Context, msg *message) error {
	h.Lock()
	for s := range h.subs {
		s.publish(ctx, msg)
	}
	h.Unlock()

	return nil
}

func (h *hub) subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{}
	h.Unlock()

	go func() {
		select {
		case <-ctx.Done():
			h.Lock()
			delete(h.subs, s)
			h.Unlock()
		}
	}()

	go s.run(ctx)

	return nil
}

func (h *hub) Subscribers() int {
	h.Lock()
	c := len(h.subs)
	h.Unlock()
	return c
}

func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}
