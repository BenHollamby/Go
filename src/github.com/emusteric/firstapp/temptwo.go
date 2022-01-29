package main

import (
	"fmt"
)

func main() {

	card := "three"
	x := "ace"
	y := "king"
	value := ParseCard(card)
	fmt.Println(value)
	sec := IsBlackjack(x, y)
	fmt.Println(sec)
}

func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "king" || card == "queen" || card == "jack" || card == "ten":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	case card == "one":
		return 1
	default:
		return 0
	}
	panic("Please implement the ParseCard function")
}

func IsBlackjack(card1, card2 string) bool {
	switch {
	case ParseCard(card1)+ParseCard(card2) == 21:
		return true
	default:
		return false
	}
	panic("Please implement the IsBlackjack function")
}
