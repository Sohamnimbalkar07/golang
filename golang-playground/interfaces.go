package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Eat interface {
	Eat() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof! I am " + d.name
}

func (d *Dog) Eat() string {
	if d == nil {
		return "I can't eat because I'm nil!"
	}
	return "I am eating " + d.name
}

func interfaces() {

	// Example 1: Using value receiver
	var s1 Speaker

	var d1 = Dog{name: "Buddy"}

	s1 = d1
	fmt.Println(s1.Speak())

	// Example 2: Using pointer receiver
	var s2 Speaker

	var d2 *Dog

	d2 = &d1

	s2 = d2
	fmt.Println(s2.Speak())

	// Example 3: Nil pointer receiver
	var e Eat

	var d3 *Dog

	e = d3
	fmt.Println(e.Eat())

	// Example 4: Uninitialized interface
	var e2 Eat
	fmt.Println(e2.Eat()) // This will cause a runtime error
}
