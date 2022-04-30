package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestSubscriber(t *testing.T) {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	h.subscribe(ctx, sub03)

	assert.Equal(t, 3, h.subscribers())

	h.unsubscribe(ctx, sub01)
	h.unsubscribe(ctx, sub02)
	h.unsubscribe(ctx, sub03)

	assert.Equal(t, 0, h.subscribers())
}

func TestCancelSubscriber(t *testing.T) {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	ctx03, cancel := context.WithCancel(ctx)
	h.subscribe(ctx03, sub03)

	assert.Equal(t, 3, h.subscribers())

	// cancel subscriber 03
	cancel()
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 2, h.subscribers())

	h.unsubscribe(ctx, sub01)
	h.unsubscribe(ctx, sub02)

	assert.Equal(t, 0, h.subscribers())
}
