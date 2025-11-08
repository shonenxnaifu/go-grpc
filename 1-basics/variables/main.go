package main

import "fmt"

// must use "var" to declare global variables
var middleName = "Cane"

func main() {
	var age int
	var name string = "john"
	var name1 = "jane"

	count := 10
	lastName := "Smith"

	// Default Values
	// Numeric = 0
	// Boolean = false
	// String = ""
	// Pointers, slice, maps, functions, and structs = nil

	middleName = "Mayor"
	fmt.Println(middleName)
	fmt.Println(age, name, name1, count, lastName)
}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName, middleName)
}
