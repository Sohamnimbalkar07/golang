package main

import (
	"fmt"
	"sync"
)

func firstRunner(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\n I am the first runner")

}

func secondRunner(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\n I am the second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	fmt.Print("Starting the runners")
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go firstRunner(wg)
	go secondRunner(wg)

	fmt.Print("\n Waiting for all runners to finish execution")

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()

	fmt.Print("\n All runners have finished execution")
}

func WaitGroup() {
	// Launching both the runners
	execute()
}
