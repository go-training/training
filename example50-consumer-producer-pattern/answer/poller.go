package answer

import (
	"context"
	"log"
	"time"
)

func NewPoller() *Poller {
	return &Poller{
		routineGroup: newRoutineGroup(),
	}
}

type Poller struct {
	routineGroup *routineGroup
}

func (p *Poller) Poll(ctx context.Context, n int) error {
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

func (p *Poller) poll(ctx context.Context, n int) error {
	log.Printf("gorutine %02d: fetch task\n", n)
	task, err := p.fetch(ctx)
	if err != nil {
		return nil
	}

	time.Sleep(200 * time.Millisecond)
	log.Printf("gorutine %02d: execute task\n", n)
	return p.execute(ctx, task)
}

func (p *Poller) fetch(ctx context.Context) (string, error) {
	// connect database or other service
	time.Sleep(400 * time.Millisecond)
	return "foobar", nil
}

func (p *Poller) execute(ctx context.Context, task string) error {
	return nil
}
