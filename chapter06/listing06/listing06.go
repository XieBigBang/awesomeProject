package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//wg.Add(4)
	wg.Add(1)

	stick := make(chan int)

	go run(stick)

	fmt.Printf("runner 1 ready to run!!\n")
	stick <- 1

	wg.Wait()
}

func run(stick chan int) {
	//defer wg.Done()

	var newRunner int

	runner := <-stick
	fmt.Printf("runner %v get the stick!!\n", runner)

	if runner < 4 {
		newRunner = runner + 1
		fmt.Printf("runner %v ready to next round!!\n", newRunner)
		go run(stick)
	}

	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("runner %v take %v second!! \n", runner, n)

	if runner == 4 {
		fmt.Printf("the last round!! \n")
		//close(stick)
		wg.Done()
		return
	}

	fmt.Printf("runner %v swith the stick to the runner %v!! \n", runner, newRunner)
	stick <- newRunner

}
