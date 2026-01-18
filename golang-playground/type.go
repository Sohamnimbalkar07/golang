package main

import "fmt"

func typeExample() {

	type Age int

	var myAge Age

	myAge = 30

	myAge = Age(25) // Type conversion from int to Age

	fmt.Println("My age is:", myAge)
}
