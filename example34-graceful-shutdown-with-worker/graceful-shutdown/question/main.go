package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
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
			c.jobsChan <- job
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

func (c *Consumer) worker(ctx context.Context, num int) {
	log.Println("start the worker", num)
	for {
		select {
		case job := <-c.jobsChan:
			if ctx.Err() != nil {
				log.Println("get next job", job, "and close the worker", num)
				return
			}
			c.process(num, job)
		case <-ctx.Done():
			log.Println("close the worker", num)
			return
		}
	}
}

const poolSize = 2

func main() {
	stop := make(chan bool)
	// create the consumer
	consumer := Consumer{
		inputChan: make(chan int, 10),
		jobsChan:  make(chan int, poolSize),
	}

	ctx := withContextFunc(context.Background(), func() {
		log.Println("cancel from ctrl+c event")
		stop <- true
	})

	for i := 0; i < poolSize; i++ {
		go consumer.worker(ctx, i)
	}

	go consumer.startConsumer(ctx)

	go func() {
		consumer.queue(1)
		consumer.queue(2)
		consumer.queue(3)
		consumer.queue(4)
		consumer.queue(5)
	}()

	<-stop
	// time.Sleep(10 * time.Second)
}
