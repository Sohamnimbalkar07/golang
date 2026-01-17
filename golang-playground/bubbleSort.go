package main

import "fmt"

func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func bubbleSort() {
	var n int
	fmt.Print("Enter number of elements (max 10): ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the numbers:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	BubbleSort(arr)

	fmt.Println("Sorted order:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
