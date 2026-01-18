package main

func applyIt(afunc func(int) int, value int) int {
	return afunc(value)
}
func functionsAsArgument() {

	square := func(x int) int {
		return x * x
	}

	result := applyIt(square, 5)

	//anonymous function as argument
	sum := applyIt(func(x int) int {
		return x + x
	}, 10)

	println(result)
	println(sum)
}
