package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Consumer struct
type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func (c *Consumer) queue(input int) {
	c.jobsChan <- input
	fmt.Println("already send input value:", input)
}

func (c *Consumer) worker(num int) {
	for job := range c.jobsChan {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		fmt.Printf("Sleeping %d seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Println("worker:", num, " job value:", job)
	}
}

const poolSize = 1

func main() {
	// create the consumer
	consumer := Consumer{
		inputChan: make(chan int, 1),
		jobsChan:  make(chan int, poolSize),
	}

	for i := 0; i < poolSize; i++ {
		go consumer.worker(i)
	}

	consumer.queue(1)
	consumer.queue(2)
	consumer.queue(3)
	consumer.queue(4)

	time.Sleep(5 * time.Second)
}
