package main

import "fmt"

func main() {
	statement, total := sum("the sum of 1,2,3 is: ", 1, 2, 3)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence3, total3 := sum("3", numbers...)
	fmt.Println("sequence: ", sequence3, "total", total3)
}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
