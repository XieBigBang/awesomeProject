package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int64

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	go inc()
	go inc()

	wg.Wait()

	fmt.Printf("%v ", count)
}

func inc() {
	defer wg.Done()

	//for i := 0; i < 2; i++ {
	//	c := count
	//	c++
	//	runtime.Gosched()
	//	count = c
	//}

	//for i := 0; i< 2; i++ {
	//	atomic.AddInt64(&count, 1)
	//	runtime.Gosched()
	//}

	for i := 0; i < 2; i++ {

		mutex.Lock()
		{
			c := count
			c++
			runtime.Gosched()
			count = c
		}
		mutex.Unlock()
	}

}
