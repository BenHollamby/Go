package main

import "fmt"

func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("neither one nor two")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
	case i <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than twenty")
	}

	x := 9
	switch {
	case x <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough //watch carefully
	case x <= 20:
		fmt.Println("less than or equal to 20")
	default:
		fmt.Println("greater than twenty")
	}
}
