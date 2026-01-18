package main

import "fmt"

type Animals interface {
	Speak()
}

type Dogs struct {
	name string
}

func (Dogs) Speak() {
	fmt.Println("Woof")
}

type Cats struct{}

func (Cats) Speak() {
	fmt.Println("Meow")
}

func interface1() {
	var a Animals

	a = Dogs{name: "Buddy"}
	a.Speak()

	// Disambiguation using type assertion
	d, ok := a.(Dogs)

	fmt.Println("Type assertion result:", d, ok)
	if ok {
		fmt.Println("This is a Dogs:", d)
	}
}
