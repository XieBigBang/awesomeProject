package main

import (
	"awesomeProject/chapter07/listing03"
	"log"
	"sync"
	"sync/atomic"
)

const GR_SIZE = 2
const JOB_SIZE = 20

func main() {

	worker, err := listing03.New(GR_SIZE)
	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(JOB_SIZE)

	for i := 0; i < JOB_SIZE; i++ {
		go func() {
			worker.Add(createTask())
			wg.Done()
		}()
	}

	wg.Wait()

	worker.Shutdown()

}

type customerTask struct {
	ID int64
}

func (c *customerTask) Work() {
	log.Printf("task %v doing work.", c.ID)
}

var id int64

func createTask() *customerTask {
	return &customerTask{atomic.AddInt64(&id, 1)}
}
