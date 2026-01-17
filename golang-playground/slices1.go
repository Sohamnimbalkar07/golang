package main

import "fmt"

func slices1() {

	//it creates an underlying array with size equal to capacity
	var slice []int = make([]int, 3)

	var slice1 []int = make([]int, 2, 5)

	slice[0] = 10
	slice[1] = 20
	// slice[2] = 30

	fmt.Println("Slice:", slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))

	fmt.Println("Slice1:", slice1)
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

	var slice2 []int = make([]int, 4)

	// capacity is 4, but as we add 5 element, new underlying array is created and slice2 now points to new array
	slice2 = []int{1, 2, 3, 4, 5}

	// append modify same underlying array
	slice2 = append(slice2, 9, 10)

	slice2 = append(slice2, 6, 7, 8)

	fmt.Println("Slice2:", slice2)
}
