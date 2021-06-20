package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal)
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
	jobChan := make(chan int, 100)
	stopped := make(chan struct{})
	finished := make(chan struct{})
	wg := &sync.WaitGroup{}
	ctx := withContextFunc(
		context.Background(),
		func() {
			log.Println("stop the server")
			close(stopped)
			wg.Wait()
			close(finished)
		},
	)

	// create 4 workers to process job
	for i := 0; i < 4; i++ {
		go func(i int) {
			log.Printf("start worker: %02d", i)
			for {
				select {
				case <-finished:
					log.Printf("stop worker: %02d", i)
					return
				default:
					select {
					case job := <-jobChan:
						time.Sleep(time.Duration(job*100) * time.Millisecond)
						log.Printf("worker: %02d, process job: %02d", i, job)
						wg.Done()
					default:
						log.Printf("worker: %02d, no job", i)
						time.Sleep(1 * time.Second)
					}
				}
			}
		}(i + 1)
	}

	// send job
	go func() {
		for i := 0; i < 40; i++ {
			wg.Add(1)
			select {
			case jobChan <- i:
				time.Sleep(100 * time.Millisecond)
				log.Printf("send the job: %02d\n", i)
			case <-stopped:
				wg.Done()
				log.Println("stoped send the job")
				return
			}
		}
		return
	}()

	select {
	case <-ctx.Done():
		log.Println("server down")
	}
}
