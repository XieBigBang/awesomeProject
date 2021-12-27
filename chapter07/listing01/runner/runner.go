package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var TimeoutErr = errors.New("timeout")
var InteruptErr = errors.New("interrupt")

type runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

func New(d time.Duration) *runner {
	return &runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *runner) Add(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

func (r *runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case e := <-r.complete:
		return e
	case <-r.timeout:
		return TimeoutErr
	}
}

func (r *runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return InteruptErr
		}
		task(id)

		// 异常处理？？？
	}
	return nil
}

func (r *runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
