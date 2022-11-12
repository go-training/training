package main

import (
	"context"

	"example/issue"
)

func main() {
	producer := issue.NewPoller(5)
	producer.Poll(context.Background())
}
