package main

import "fmt"

func slices() {

	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var s1 []int = arr[1:4] // Slicing from index 1 to 3

	var s2 []int = arr[:] // Slicing the entire array

	fmt.Println("Array:", arr)
	fmt.Println("Sliced Slice:", s1)
	fmt.Println("Sliced Slice (entire array):", s2)

	// Modifying the slice
	s1[0] = 20
	fmt.Println("Array:", arr)
	fmt.Println("Sliced Slice:", s1)

	fmt.Println("Length of s1:", len(s1))
	fmt.Println("Capacity of s1:", cap(s1))

}
