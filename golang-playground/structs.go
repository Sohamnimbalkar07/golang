package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func main() {

	var p1 Person

	p1.name = "Soham"

	fmt.Println("name of person is", p1.name)

	var p2 = new(Person)

	fmt.Println("P2", p2)

	p3 := Person{name: "Soham", age: 25, address: "Pune"}

	fmt.Println("p3", p3)
}
