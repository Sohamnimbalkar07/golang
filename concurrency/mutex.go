package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func inc() {
	mu.Lock()
	i = i + 1
	mu.Unlock()
	wg.Done()
}

func mutex() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
