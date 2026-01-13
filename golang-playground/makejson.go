package main

import (
	"encoding/json"
	"fmt"
)

func makeJSON() {
	var name string
	var address string

	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	data := make(map[string]string)
	data["name"] = name
	data["address"] = address

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))
}
