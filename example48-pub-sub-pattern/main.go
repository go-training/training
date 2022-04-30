package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	h.subscribe(ctx, sub03)

	_ = h.publish(ctx, &message{data: []byte("test01")})
	_ = h.publish(ctx, &message{data: []byte("test02")})
	_ = h.publish(ctx, &message{data: []byte("test03")})
	time.Sleep(1 * time.Second)

	h.unsubscribe(ctx, sub03)
	_ = h.publish(ctx, &message{data: []byte("test04")})
	_ = h.publish(ctx, &message{data: []byte("test05")})

	time.Sleep(1 * time.Second)
}
