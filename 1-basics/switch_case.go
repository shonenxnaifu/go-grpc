package main

import (
	"fmt"
)

func main() {
	fruit := "pineapple"

	switch fruit {
	case "apple":
		fmt.Println("its an apple")
	case "banana":
		fmt.Println("its an apple")
	default:
		fmt.Println("unkown fruit")
	}

	// Multiple conditions
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("its a weekday")
	case "Sunday":
		fmt.Println("its a weekend")
	default:
		fmt.Println("Invalid day")
	}

	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not 2")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType("true")
}

// we cant use fallthrough in type assertion / type switch
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("its an integer")
	case float64:
		fmt.Println("its a float")
	case string:
		fmt.Println("its a string")
	default:
		fmt.Println("unkown type")
	}
}

