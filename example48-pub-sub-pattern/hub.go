package main

import (
	"context"
	"sync"
)

type hub struct {
	sync.Mutex
	subs map[*subscriber]struct{}
}

func (h *hub) subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{}
	h.Unlock()

	go s.run(ctx)

	return nil
}

func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}
