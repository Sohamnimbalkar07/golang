package main

import (
	"fmt"
	"time"
)

func abort() {
	abort := make(chan bool)

	go func() {
		for {
			select {
			case <-abort:
				fmt.Println("aborted")
				return
			default:
				fmt.Println("working")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	abort <- false // abort signal (send, not close)

	time.Sleep(1 * time.Second)
}
