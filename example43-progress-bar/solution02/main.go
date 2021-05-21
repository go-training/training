package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	var limit int64 = 1024 * 1024 * 10000
	// we will copy 10 Gb from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	// start new bar
	bar := pb.Full.Start64(limit)
	finishCh := make(chan struct{})
	go func(bar *pb.ProgressBar) {
		d := time.NewTicker(2 * time.Second)
		startTime := bar.StartTime()
		// Using for loop
		for {
			// Select statement
			select {
			case <-finishCh:
				d.Stop()
				log.Println("finished")
				return
			// Case to print current time
			case <-d.C:
				if !bar.IsStarted() {
					continue
				}
				currentTime := time.Now()
				dur := currentTime.Sub(startTime)
				lastSpeed := float64(bar.Current()) / dur.Seconds()
				remain := float64(bar.Total() - bar.Current())
				remainDur := time.Duration(remain/lastSpeed) * time.Second
				fmt.Println("Progress:", float32(bar.Current())/float32(bar.Total())*100)
				fmt.Println("last speed:", lastSpeed/1024/1024)
				fmt.Println("remain suration:", remainDur)
			}
		}
	}(bar)
	// create proxy reader
	barReader := bar.NewProxyReader(reader)
	// copy from proxy reader
	if _, err := io.Copy(writer, barReader); err != nil {
		log.Fatal(err)
	}
	// finish bar
	bar.Finish()
	close(finishCh)
}
