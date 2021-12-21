package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("begin..........")
	go findPrime("A")
	go findPrime("B")

	fmt.Println("wait..........")
	wg.Wait()
	fmt.Println("end..........")

}

func findPrime(pre string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%v:%v\n", pre, outer)
	}
}
