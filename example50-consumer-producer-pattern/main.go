package main

import (
	"context"

	// poller02 "example/answer"
	poller01 "example/issue"
)

func main() {
	// issue
	producer01 := poller01.NewPoller(5)
	producer01.Poll(context.Background())

	// answer
	// producer02 := poller02.NewPoller(5)
	// producer02.Poll(context.Background())
}
