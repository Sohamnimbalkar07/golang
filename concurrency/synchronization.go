package main

import (
	"fmt"
	"time"
)

var x int = 1
var done = make(chan bool)

func Task1() {
	x = x + 1
	done <- true
}

func Task2() {
	<-done
	fmt.Printf("value of x: %d", x)
}

func synchronization() {
	go Task1()
	go Task2()

	time.Sleep(100 * time.Millisecond)
}
