package main

import "fmt"

func pointers() {

	var p *int

	var a int = 10

	p = &a

	// modify value using pointer
	*p = 15

	fmt.Println("value of a is ", *p)

}
