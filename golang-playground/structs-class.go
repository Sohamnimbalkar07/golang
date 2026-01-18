package main

import "fmt"

type Car struct {
	Brand string
	Speed int
}

func (c Car) Display() {
	fmt.Println("Brand:", c.Brand)
	fmt.Println("Speed:", c.Speed)
}

func (c Car) Accelerate() Car {
	c.Speed += 10
	return c
}

func (c *Car) setBrand(brand string) {
	c.Brand = brand
}

func tests() {

	car := Car{
		Brand: "Tesla",
		Speed: 50,
	}

	car.Display()
	var newCar = car.Accelerate()
	car.setBrand("BMW")
	car.Display()
	newCar.Display()
}
