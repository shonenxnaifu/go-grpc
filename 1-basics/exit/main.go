package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("deferred statement")

	fmt.Println("starting the main function")

	// exit with status code of 1
	os.Exit(1)

	//this will never be executed
	fmt.Println("end of main function")
}
