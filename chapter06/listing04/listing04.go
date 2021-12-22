package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var shutdow int64

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(2000 * time.Millisecond)

	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdow, 1)

	wg.Wait()

}

func doWork(worker string) {
	defer wg.Done()
	for {
		fmt.Printf("doing %v work \n", worker)
		time.Sleep(1000 * time.Millisecond)

		if atomic.LoadInt64(&shutdow) == 1 {
			fmt.Printf("%v stop \n", worker)
			break
		}
	}
}
