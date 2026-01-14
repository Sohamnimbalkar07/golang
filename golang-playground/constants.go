package main

import "fmt"

func constants() {

	fmt.Println("Constants")

	const a = 10

	const pi float32 = 3.14

	const b = a

	// 11, 22, 33
	const (
		x = 11
		y = 22
		z = 33
	)

	// 0 1 2
	const (
		low = iota
		medium
		high
	)

	// 2 3
	const (
		p = iota + 2
		q
	)

	type Status int

	//0 1 2
	const (
		Active Status = iota
		Inactive
		Suspended
	)

	fmt.Println(pi, a, b, x, y, z, low, medium, high, p, q)

	fmt.Println(Active, Inactive, Suspended)

}
