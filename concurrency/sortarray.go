package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortPart(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting subarray:", arr)
	sort.Ints(arr)
}

func sortarray() {
	var nums []int
	fmt.Println("Enter integers (end with non-number):")

	for {
		var x int
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		nums = append(nums, x)
	}

	n := len(nums)
	if n == 0 {
		fmt.Println("No numbers entered")
		return
	}

	partSize := n / 4
	var wg sync.WaitGroup

	// Create 4 partitions
	parts := make([][]int, 4)
	for i := 0; i < 4; i++ {
		start := i * partSize
		end := start + partSize
		if i == 3 {
			end = n // last part takes remaining elements
		}
		parts[i] = nums[start:end]
	}

	// Sort each partition concurrently
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sortPart(parts[i], &wg)
	}

	wg.Wait()

	// Merge all parts
	result := append(parts[0], parts[1]...)
	result = append(result, parts[2]...)
	result = append(result, parts[3]...)
	sort.Ints(result)

	fmt.Println("Final sorted array:", result)
}
