package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wgr sync.WaitGroup

func initApp() {
	fmt.Println("initialized")
}

func doStuff() {
	once.Do(initApp)

	fmt.Println("Hello")
	wgr.Done()
}

func syncOnce() {
	wgr.Add(2)
	go doStuff()
	go doStuff()
	wgr.Wait()
}
