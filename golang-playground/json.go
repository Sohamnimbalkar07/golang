package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Address string
}

func json1() {

	s := Student{Name: "Soham", Age: 25, Address: "Pune"}

	var s1 Student

	fmt.Println("Student:", s)

	var data, err = json.Marshal(s)

	fmt.Println(data, err)

	json.Unmarshal(data, &s1)

	fmt.Println("Student s1", s1)

}
