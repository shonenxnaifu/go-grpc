package main

import "fmt"

func main()  {
	process(10)
}

func process(i int) {
	defer fmt.Println("deferred i value: ", i)
	defer fmt.Println("first deferred last")
	defer fmt.Println("second deferred 2nd")
	defer fmt.Println("third deferred 1st")
	i++
	fmt.Println("normal")
	fmt.Println("value of i: ", i)
}
