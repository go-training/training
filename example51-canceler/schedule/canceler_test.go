package schedule

import (
	"context"
	"testing"
	"time"
)

func TestUserCancelTask(t *testing.T) {
	var canceled bool
	var err error
	engine := newCanceler()
	stop := make(chan struct{})
	go func() {
		canceled, err = engine.Cancelled(context.Background(), "test123456")
		stop <- struct{}{}
	}()
	time.Sleep(1 * time.Second)
	engine.Cancel(context.Background(), "test123456")
	<-stop

	if !canceled {
		t.Fatal("can't cancel task")
	}
	if err != nil {
		t.Fatal("get error")
	}
}

func TestContextCancelTask(t *testing.T) {
	var canceled bool
	var err error
	engine := newCanceler()
	stop := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		canceled, err = engine.Cancelled(ctx, "test123456")
		stop <- struct{}{}
	}()
	cancel()
	<-stop

	if canceled {
		t.Fatal("detect cancel task")
	}
	if err != nil {
		t.Fatal("get error")
	}
}

func TestUserCancelTaskFirst(t *testing.T) {
	var canceled bool
	var err error
	engine := newCanceler()

	// User cancel task first and
	_ = engine.Cancel(context.Background(), "test1234")
	canceled, err = engine.Cancelled(context.Background(), "test1234")

	if !canceled {
		t.Fatal("can't get cancel event")
	}
	if err != nil {
		t.Fatal("get error")
	}
}
