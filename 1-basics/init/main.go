package main

import "fmt"

func init() {
	fmt.Println("initializing package1...")
}

func init() {
	fmt.Println("initializing package2...")
}

func init() {
	fmt.Println("initializing package3...")
}

func main () {
	fmt.Println("inside the main function")
}
