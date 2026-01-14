package main

import "fmt"

func loops() {

	x := 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is smaller than 5")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i", i)
	}

	j := 0

	for j < 5 {
		fmt.Println("value of j", j)
		j++
	}

	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}
}
