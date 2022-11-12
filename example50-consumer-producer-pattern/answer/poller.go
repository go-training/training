package answer

import (
	"context"
	"log"
	"sync"
	"time"
)

func NewPoller() *Poller {
	return &Poller{
		routineGroup: newRoutineGroup(),
		ready:        make(chan struct{}, 1),
		metric:       newMetric(),
	}
}

type Poller struct {
	sync.Mutex
	ready        chan struct{}
	routineGroup *routineGroup
	metric       *metric
}

func (p *Poller) schedule(n int) {
	p.Lock()
	defer p.Unlock()
	if int(p.metric.BusyWorkers()) >= n {
		return
	}

	select {
	case p.ready <- struct{}{}:
	default:
	}
}

func (p *Poller) Poll(ctx context.Context, n int) error {
	// scheduler
	for {
		p.schedule(n)

		select {
		case <-p.ready:
		case <-ctx.Done():
			return nil
		}

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				task, err := p.fetch(ctx)
				if err != nil {
					log.Println("fetch task error:", err.Error())
					break
				}
				p.metric.IncBusyWorker()
				p.routineGroup.Run(func() {
					if err := p.execute(ctx, task); err != nil {
						log.Println("execute task error:", err.Error())
					}
				})
				break LOOP
			}
		}
	}
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
	defer func() {
		p.metric.DecBusyWorker()
	}()
	return nil
}
