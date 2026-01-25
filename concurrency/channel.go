package main

import "fmt"

func prod(i1 int, i2 int, ch chan int) {
	prod := i1 * i2
	ch <- prod
}
func main() {

	ch := make(chan int)
	go prod(2, 3, ch)
	go prod(4, 5, ch)
	a := <-ch
	b := <-ch
	result := a * b
	fmt.Println("Product is:", result)
}
