package main

import (
	"fmt"
	"os"
)

// os.Args slice contains the command-line arguments passed to the program.
// os.Args[0] is path where built binary is located
// os.Args[1:] are the actual arguments passed to the program
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Hello,", os.Args[0])
	} else {
		fmt.Println("Hello, World!")
	}
	fmt.Println("first os Args", os.Args[0])
}
