package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	round := make(chan int)

	go play("A", round)
	go play("B", round)

	fmt.Printf("game begin..........")
	time.Sleep(1 * time.Second)
	round <- 1

	wg.Wait()
}

func play(player string, rr chan int) {
	defer wg.Done()

	fmt.Printf("player %v ready now!\n", player)
	for {
		ball, ok := <-rr
		if !ok {
			fmt.Printf("player %v win!\n", player)
			return
		}

		n := rand.Intn(100)
		//fmt.Printf("------------%v--------\n", n)
		if n%13 == 0 {
			fmt.Printf("play %v lose the ball!\n", player)
			close(rr)
			return
		}
		ball++
		fmt.Printf("player %v get the ball %v round\n", player, ball)

		rr <- ball
	}

}
