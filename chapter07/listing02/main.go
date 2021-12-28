package main

import (
	"awesomeProject/chapter07/listing02/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type connection struct {
	Cid int64
}

func (c *connection) Close() error {
	log.Println("connection close")
	return nil
}

var snum int64

func createCon() (io.Closer, error) {
	id := atomic.AddInt64(&snum, 1)
	return &connection{id}, nil
}

const QNUM = 20

var wg sync.WaitGroup

func main() {

	wg.Add(QNUM)

	p, err := pool.New(createCon, 3)

	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < QNUM; i++ {

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		go func(q int) {
			doQuery(q, *p)
			wg.Done()
		}(i)

	}
	wg.Wait()

	p.Close()

}

func doQuery(id int, p pool.Pool) {
	c, err := p.Aquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(c)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("do query connection id:%v,  qid:%v", c.(*connection).Cid, id)

}
