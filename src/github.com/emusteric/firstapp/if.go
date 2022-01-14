package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	number := 50
	guess := 40

	if guess < 1 || guess > 100 { //OR operand
		fmt.Println("The guess must be between 1 and 100")
	}
	if guess >= 1 && guess <= 100 { //AND operand
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("Exactly")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	number1 := 50
	guess1 := 60

	if guess1 < 1 || guess1 > 100 { //OR operand
		fmt.Println("The guess must be between 1 and 100")
	} else {
		if guess1 < number1 {
			fmt.Println("Too low")
		}
		if guess1 > number1 {
			fmt.Println("Too high")
		}
		if guess1 == number1 {
			fmt.Println("Exactly")
		}
		fmt.Println(number1 <= guess1, number1 >= guess1, number1 != guess1)
	}

}
