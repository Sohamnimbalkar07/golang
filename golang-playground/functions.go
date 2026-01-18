package main

import "fmt"

// Closure example: function returning another function
func addBy(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// variable argument function
func getMax(vals ...int) int {
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func functions() {

	addFive := addBy(5)

	result := addFive(10)

	max := getMax(1, 2, 3, 4, 5)

	vslice := []int{10, 20, 5, 30, 15}
	maximum := getMax(vslice...)

	fmt.Println(result)
	fmt.Println(max)
	fmt.Println(maximum)
}
