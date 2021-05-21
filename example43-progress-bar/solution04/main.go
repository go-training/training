package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cheggaaa/pb/v3"
)

type readerFunc func(p []byte) (n int, err error)

func (rf readerFunc) Read(p []byte) (n int, err error) { return rf(p) }

func copy(ctx context.Context, dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, readerFunc(func(p []byte) (int, error) {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			return src.Read(p)
		}
	}))
	return err
}

func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			f()
			cancel()
		}
	}()

	return ctx
}

func main() {

	ctx := withContextFunc(
		context.Background(),
		func() {
			// clear machine field
			log.Println("interrupt received, terminating process")
		},
	)

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
			case <-ctx.Done():
				d.Stop()
				log.Println("interrupt received")
				return
			case <-finishCh:
				d.Stop()
				log.Println("finished")
				return
			// Case to print current time
			case <-d.C:
				if ctx.Err() != nil {
					return
				}
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
	if err := copy(ctx, writer, barReader); err != nil {
		log.Println("cancel upload data:", err.Error())
	}
	// finish bar
	bar.Finish()
	close(finishCh)
	time.Sleep(1 * time.Second)
}
