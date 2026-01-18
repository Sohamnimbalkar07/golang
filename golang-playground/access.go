package main

import (
	"fmt"
	data "golang-playground/access"
)

func access() {
	fmt.Println("Hello, World!")

	// fmt.Println("i is ", data.i) //error as i is unexported variable because it starts with lowercase letter

	fmt.Println("Str is", data.Str) // prints Initial Name

	data.ModifyName()

	// data.modifyI() //error as modifyI is unexported function because it starts with lowercase letter

	fmt.Println("After modification, Str is", data.Str) // prints Modified Name

	p := data.Point{X: 10, Y: 20}

	fmt.Println("Point is", p) // prints {10 20}

	// no need to reference p as Go does it automatically ( referencing = &p.Move(5, -5) )
	p.Move(5, -5)
	fmt.Println("After moving, Point is", p) // prints {15 15}

}
