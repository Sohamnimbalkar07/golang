package main

import (
	"fmt"
	"time"
)

func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "data from channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "data from channel 2"
	}()
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)

	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}
