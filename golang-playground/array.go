package main

import "fmt"

func array() {

	var arr [5]int

	arr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	arr2 := [3]string{"Soham", "is", "here"}

	fmt.Println(arr2)

	var arr3 [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println(arr3)

	arr4 := [...]int{11, 22, 33, 44, 55}

	fmt.Println(arr4)
}
