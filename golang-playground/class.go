package main

import "fmt"

type myInt int

func (i myInt) Double() int {
	return int(i) * 2
}

func class() {

	var num myInt = 10
	fmt.Println("Doubled number:", num.Double()) //20
	fmt.Printf("Type of num: %T\n", num)         // main.myInt
	fmt.Println("Value of num:", num)            //10 ( functions are call by value so original value remains unchanged)

}
