package main

import (
	"fmt"
	"sync"
	"time"
)

type Foo struct {
	Value int
	Mutex sync.Mutex
}

func (f *Foo) Inc(wg *sync.WaitGroup) {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()
	value := f.Value + 1
	time.Sleep(time.Millisecond)
	f.Value = value
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	f := Foo{Value: 0}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go f.Inc(&wg)
	}
	wg.Wait()
	fmt.Println(f.Value)
}
