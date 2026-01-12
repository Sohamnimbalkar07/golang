package main

import (
	"fmt"
	"sort"
)

func sliceExample() {
	s := make([]int, 0, 3)

	for {
		var input string
		fmt.Print("Enter number X to quit")
		fmt.Scan(&input)

		if input == "X" || input == "x" {
			break
		}

		var num int
		fmt.Sscanf(input, "%d", &num)

		s = append(s, num)
		sort.Ints(s)

		fmt.Println(s)
	}
}
