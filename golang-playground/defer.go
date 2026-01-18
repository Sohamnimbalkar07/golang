package main

import "fmt"

func main() {

	defer fmt.Println("Deferred: This will be printed last.")

	fmt.Println("This will be printed first.")

	i := 1

	defer fmt.Println("Deferred: The value of i is", i)

	i = 2
	fmt.Println("The value of i is", i)
}
