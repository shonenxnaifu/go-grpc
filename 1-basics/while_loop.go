package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for as while with break
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum: ", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	// num := 1
	// for num <= 10 {
	// 	if num%2 == 0 {
	// 		num++
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number: ", num)
	// 	num++
	// }

	// guessing game
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100)

	//welcome image
	fmt.Println("welcome to the guessing game!")
	fmt.Println("i have choosen number between 1 and 100")
	fmt.Println("can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		//check if the guess if correct
		if guess == target {
			fmt.Println("Congratulation! you guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("too low! try guessing a higher number")
		} else {
			fmt.Println("too high! try guessing a lower number")
		}
	}
}
