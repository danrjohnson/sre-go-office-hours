package main

import (
	"fmt"
	"sync"
	"time"
)

func TheUltimatePingPonger(receiver chan int, sender chan int, foo string) {
	for {
		_ = <-receiver
		fmt.Println(foo)
		sender <- 90
		time.Sleep(time.Duration(time.Second * 1))
	}
}

func main() {
	pinger := make(chan int)
	ponger := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go TheUltimatePingPonger(pinger, ponger, "Ping")
	go TheUltimatePingPonger(ponger, pinger, "Pong")
	pinger <- 420
	wg.Wait()
}
