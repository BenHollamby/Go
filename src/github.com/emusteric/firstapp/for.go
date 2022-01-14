package main

import (
	"fmt"
)

func main() {
	var x int = 0
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for ; x < 5; x++ {
	}
	fmt.Println("\n", x)

	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}

	z := 0
	for {
		fmt.Println(z)
		z++
		if z == 10 {
			break
		}
	}

	for q := 0; q < 10; q++ {
		if q%2 == 0 {
			continue
		}
		fmt.Println(q)
	}

Loop:
	for b := 1; b <= 3; b++ {
		for h := 1; h <= 3; h++ {
			fmt.Println(b * h)
			if b*h >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}

}
