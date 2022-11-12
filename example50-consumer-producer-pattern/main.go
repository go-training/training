package main

import (
	"context"

	"example/issue"
)

func main() {
	producer := issue.NewPoller()
	producer.Poll(context.Background(), 10)
}
