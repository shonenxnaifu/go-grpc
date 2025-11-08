package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println(sum)

	greet := func() {
		fmt.Println("hello anonymous function")
	}

	greet()

	operation := add
	result := operation(3, 5)
	fmt.Println(result)

	// passing a function as an argument
	result2 := applyOperation(5, 3, add)
	fmt.Println("apply operation: ", result2)

	//returning and using a function
	multiplyFunc := createMultiplier(3)
	fmt.Println("multiply: ", multiplyFunc(6))
}

func add(a, b int) int {
	return a + b
}

// function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
