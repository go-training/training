package main

import (
	"fmt"
	"time"
)

type logEntry struct {
	message string
	time    time.Time
}

func main() {
	logCh := make(chan logEntry, 100)
	go logger(logCh)
	logCh <- logEntry{"App Start", time.Now()}
	logCh <- logEntry{"App End", time.Now()}
	time.Sleep(100 * time.Millisecond)
}

func logger(logCh <-chan logEntry) {
	for v := range logCh {
		fmt.Printf("%v: %v\n", v.time.Format("2006-01-02T01:01:01"), v.message)
	}
}
