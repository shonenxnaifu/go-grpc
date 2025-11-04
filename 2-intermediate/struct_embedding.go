package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	empInfo person // Embedded struct Named field
	// person // Anonymous field
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and i earn %2.f.\n", e.empInfo.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		empInfo: person{name: "John", age: 30},
		empId:   "E001",
		salary:  50000,
	}

	fmt.Println("Name: ", emp.empInfo.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age: ", emp.empInfo.age)   // Same as above
	fmt.Println("Emp ID: ", emp.empId)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
}
