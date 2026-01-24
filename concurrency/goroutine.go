package main

import (
	"fmt"
	"time"
)

// This program has a race condition because two goroutines concurrently access
// and modify the shared variable counter without synchronization. The statement
// counter = counter + 1 is not atomic, so the read and write operations from
// different goroutines can interleave. As a result, some increments are lost,
// and the final value of counter becomes unpredictable.

// Shared variable accessed by multiple goroutines
var counter int = 0

// increment() increases the value of counter
func increment() {
	for i := 0; i < 1000; i++ {

		// RACE CONDITION HERE ↓
		// counter = counter + 1 is NOT an atomic operation.
		// It actually does:
		// 1. Read counter
		// 2. Add 1
		// 3. Write back to counter
		//
		// If two goroutines execute this at the same time,
		// their operations can interleave, causing lost updates.
		counter = counter + 1
	}
}

func main() {

	// Start first goroutine
	go increment()

	// Start second goroutine
	go increment()

	// Wait so that both goroutines get time to execute
	// (This is NOT proper synchronization, just for simplicity)
	time.Sleep(1 * time.Second)

	// Expected value: 2000
	// Actual value: often less than 2000 due to race condition
	fmt.Println("Final counter value:", counter)
}
