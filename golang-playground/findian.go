package main

import (
	"fmt"
	"strings"
)

func findian() {

	var name string
	fmt.Print("Enter your full name: ")
	fmt.Scanln(&name)

	s := strings.ToLower(name)

	if strings.HasPrefix(s, "i") &&
		strings.HasSuffix(s, "n") &&
		strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
