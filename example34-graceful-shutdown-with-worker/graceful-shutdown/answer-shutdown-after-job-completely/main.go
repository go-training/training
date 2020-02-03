package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Consumer struct
type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func getRandomTime() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
			f()
		}
	}()

	return ctx
}

func (c *Consumer) queue(input int) bool {
	select {
	case c.inputChan <- input:
		log.Println("already send input value:", input)
		return true
	default:
		return false
	}
}

func (c Consumer) startConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.inputChan:
			if ctx.Err() != nil {
				close(c.jobsChan)
				return
			}

			select {
			case c.jobsChan <- job:
			default:
				log.Println("job channel has been closed. num:", job)
			}
		case <-ctx.Done():
			close(c.jobsChan)
			return
		}
	}
}

func (c *Consumer) process(num, job int) {
	n := getRandomTime()
	log.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	log.Println("worker:", num, " job value:", job)
}

func (c *Consumer) worker(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("start the worker", num)

	for job := range c.jobsChan {
		c.process(num, job)
	}
}

const poolSize = 5

func main() {
	finished := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(poolSize)
	// create the consumer
	consumer := Consumer{
		inputChan: make(chan int, 10),
		jobsChan:  make(chan int, poolSize),
	}

	ctx := withContextFunc(context.Background(), func() {
		log.Println("cancel from ctrl+c event")
		wg.Wait()
		close(finished)
	})

	for i := 0; i < poolSize; i++ {
		go consumer.worker(i, wg)
	}

	go consumer.startConsumer(ctx)

	go func() {
		consumer.queue(1)
		consumer.queue(2)
		consumer.queue(3)
		consumer.queue(4)
		consumer.queue(5)
		consumer.queue(6)
		consumer.queue(7)
		consumer.queue(8)
		consumer.queue(9)
		consumer.queue(10)
	}()

	<-finished
	log.Println("Game over")
}
