package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const worker = 4
const jobs = 50

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg.Add(worker)

	jq := make(chan int, 10)

	for i := 0; i < worker; i++ {
		go doWork(i, jq)
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < jobs; i++ {
		jq <- i
	}

	fmt.Println("there are no new jobs!!")
	close(jq)

	wg.Wait()
}

func doWork(worker int, jq chan int) {
	defer wg.Done()

	for {
		job, ok := <-jq
		if !ok {
			fmt.Printf("there are no new jobs!! main %v out of line!!\n", worker)
			break
		}

		fmt.Printf("main %v is doing the job %v\n", worker, job)
		n := rand.Intn(5)
		time.Sleep(time.Duration(n) * 1000 * time.Millisecond)
		fmt.Printf("main %v took %v secend finish the jos %v\n", worker, n, job)
	}

}
