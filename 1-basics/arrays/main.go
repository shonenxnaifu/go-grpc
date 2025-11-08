package main

import "fmt"

func main()  {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)
	
	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println(fruits)

	fmt.Println(fruits[2])

	originalArray := [3]int{1,2,3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("ori: ", originalArray)
	fmt.Println("copied: ", copiedArray)

	for i := 0; i<len(numbers);i++{
		fmt.Println("Element at index, ", i, ":", numbers[i])
	}

	// underscore is blank indetinfier to store unused value
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	fmt.Println("the length: ", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1,2,3}
	array2 := [3]int{10,2,3}

	fmt.Println("array1 is equal to array2: ", array1 == array2)

	var matrix[3][3]int = [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println("mtarix:", matrix)
	fmt.Println("mtarix:", matrix[1][1])
	fmt.Println("mtarix:", matrix[2][2])

	originalArray2 := [3]int{1,2,3}
	var copiedArray2 *[3]int
	copiedArray2 = &originalArray2
	copiedArray2[0] = 100

	fmt.Println("ori: ", originalArray2)
	fmt.Println("copied: ", *copiedArray2)
}

func someFunction()(int, int) {
	return 1,2
}
