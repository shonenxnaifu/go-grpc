package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase{Structs, Interfaces, Enums}

	// snake_case{variables, constants, file_names}

	// UPPERCASE{constants}

	// mixedCase{varriables, identifiers}

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
