// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create
// a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been r
// ead from the file, your program will have a slice containing one struct for each
// line in the file. After reading all lines from the file, your program should iterate
// through your slice of structs and print the first and last names found in each struct.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Name struct {
	firstName string
	lastName  string
}

func main() {
	var fileName string
	fmt.Print("Enter file name: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	names := []Name{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first, last string
		fmt.Sscan(scanner.Text(), &first, &last)

		n := Name{
			firstName: first,
			lastName:  last,
		}

		names = append(names, n)
	}

	for _, n := range names {
		fmt.Println(n.firstName, n.lastName)
	}
}
