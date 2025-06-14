package main

import (
	"sync"
)

var mux sync.Mutex

var counter int

func increment() {
	mux.Lock()
	defer mux.Unlock()
	counter++
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	println(counter)
}
