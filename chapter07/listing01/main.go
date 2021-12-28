package main

import (
	"awesomeProject/chapter07/listing01/runner"
	"log"
	"os"
	"time"
)

func main() {
	r := runner.New(30 * time.Second)
	r.Add(createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask())

	if e := r.Start(); e != nil {
		switch e {
		case runner.InteruptErr:
			log.Println("terminate interrupt")
			os.Exit(1)
		case runner.TimeoutErr:
			log.Println("terminate timeout")
			os.Exit(2)
		}
	}
	log.Printf("precess success")
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("Precess %v", i)
		time.Sleep(time.Second)
	}
}
