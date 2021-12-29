package listing03

import (
	"errors"
	"log"
	"sync"
)

type Task interface {
	Work()
}

type worker struct {
	tasks chan Task

	wg sync.WaitGroup
}

func New(size int) (*worker, error) {
	if size <= 0 {
		log.Printf("wrong param")
		return nil, errors.New("wrong size")
	}

	workers := worker{
		tasks: make(chan Task),
	}

	workers.wg.Add(size)

	for i := 0; i < size; i++ {
		go func(w int) {
			for {
				t, ok := <-workers.tasks
				if !ok {
					log.Printf("worker %v had shutdown.", w)
					break
				}
				log.Printf("worker %v going to do job.", w)
				t.Work()
			}
			workers.wg.Done()
		}(i)
	}
	return &workers, nil
}

func (w *worker) Add(t Task) {
	w.tasks <- t
}

func (w *worker) Shutdown() {
	close(w.tasks)
	w.wg.Wait()
}
