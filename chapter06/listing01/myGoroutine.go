package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start...........")

	go func() {
		defer wg.Done()

		for count := 1; count <= 3; count++ {
			for achar := 'A'; achar < 'A'+26; achar++ {
				fmt.Printf("%c ", achar)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 1; count <= 3; count++ {
			for achar := 'a'; achar < 'a'+26; achar++ {
				fmt.Printf("%c ", achar)
			}
		}
	}()

	fmt.Println("Wait............")
	wg.Wait()
	fmt.Println("End............")

}
