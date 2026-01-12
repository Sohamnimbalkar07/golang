package main

import "fmt"

func truncate() {

	var f float64
	fmt.Print("enter floating point number")
	fmt.Scan(&f)

	var truncated int
	truncated = int(f)
	fmt.Println("Truncated value is", truncated)
}
