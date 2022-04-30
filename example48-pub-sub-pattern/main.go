package main

import "context"

func main() {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	h.subscribe(ctx, sub03)
}
