package main

import "fmt"

func main()  {
	// panic(interface{})

	// example valid input
	process(10)

	//example invalid input
	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be a non-negative number")
		// fmt.Println("After panic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("processing input: ", input)
}
