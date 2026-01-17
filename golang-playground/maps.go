package main

import "fmt"

func maps() {

	var idMap map[string]int = make(map[string]int)

	idMap["a"] = 1

	fmt.Println(idMap)

	var m map[string]int = map[string]int{}

	m["b"] = 2

	fmt.Println(m)

	delete(m, "b")

	id, p := m["b"]

	// 2 true
	fmt.Println(id, p)
}
