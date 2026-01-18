package main

import "fmt"

var funcVarn func(int) int

func increment(i int) int {
	return i + 1
}

func variablesAsFunctions() {

	funcVarn = increment

	num := 5
	fmt.Printf("Value before incrementing: %d\n", num)
	num = funcVarn(num)
	fmt.Printf("Value after incrementing: %d\n", num)
}
