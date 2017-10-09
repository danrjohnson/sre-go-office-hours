package main

import (
	"fmt"
	"sync"
	"time"
)

func ThePrinter(wg *sync.WaitGroup, i int) {
	fmt.Printf("I printed %d\n", i)
	time.Sleep(time.Duration(3 * time.Second))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go ThePrinter(&wg, i)
	}
	wg.Wait()
}
