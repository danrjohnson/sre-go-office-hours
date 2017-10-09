package main

import (
	"fmt"
	"sync"
)

const (
	NOTDONE = "nope"
	DONE    = "yup"
)

func Producer(channel chan bool, wg *sync.WaitGroup) {
	channel <- true
	fmt.Println("I completed")
	wg.Done()
}

func Waiter(channel chan bool) {
	for _ = range channel {
		fmt.Println("It finished.")
	}
}

func Closer(channel chan bool) {
	close(channel)
}

func main() {
	isDone := make(chan bool, 200)
	go Producer(isDone, &wg)
	go Producer(isDone, &wg)
	go Closer(isDone)
	Waiter(isDone)
}
