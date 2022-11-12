package main

import (
	"context"
	"log"
	"time"
)

func main() {
	producer := &poller{
		routineGroup: newRoutineGroup(),
	}

	producer.Poll(context.Background(), 10)
}

type poller struct {
	routineGroup *routineGroup
}

func (p *poller) Poll(ctx context.Context, n int) error {
	for i := 0; i < n; i++ {
		func(i int) {
			p.routineGroup.Run(func() {
				for {
					select {
					case <-ctx.Done():
						return
					default:
						_ = p.poll(ctx, i)
					}
				}
			})
		}(i)
	}
	p.routineGroup.Wait()
	return nil
}

func (p *poller) poll(ctx context.Context, n int) error {
	task, err := p.fetch(ctx)
	if err != nil {
		return nil
	}

	log.Printf("gorutine %02d: execute task\n", n)
	return p.execute(ctx, task)
}

func (p *poller) fetch(ctx context.Context) (string, error) {
	// connect database or other service
	time.Sleep(400 * time.Millisecond)
	return "foobar", nil
}

func (p *poller) execute(ctx context.Context, task string) error {
	return nil
}
