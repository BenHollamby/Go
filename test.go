package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const a string = "rock"
	const b string = "paper"
	const c string = "scissors"

	var x [3]string
	x[0] = a
	x[1] = b
	x[2] = c

	random1 := rand.Intn(len(x))
	random2 := rand.Intn(len(x))
	cc1 := random1
	cc2 := random2
	fmt.Println(cc1)
	fmt.Println(cc2)

	if cc1 == cc2 {
		fmt.Println("its a tie!")
	}

	if cc1 == 0 && cc2 == 1 {
		fmt.Println("cc2 wins!")
	}

	if cc1 == 0 && cc2 == 2 {
		fmt.Println("cc1 wins!")
	}

	if cc1 == 1 && cc2 == 0 {
		fmt.Println("cc1 wins!")
	}

	if cc1 == 1 && cc2 == 2 {
		fmt.Println("cc2 wins!")
	}

	if cc1 == 2 && cc2 == 0 {
		fmt.Println("cc2 wins!")
	}

	if cc1 == 2 && cc2 == 1 {
		fmt.Println("cc1 wins!")
	}

}

/*
func zeroOne(y) {
	rand.Seed(time.Now().UnixNano())
	in := y
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	return pick
}


func computerOne() string {
	var result string
	rand.Seed(time.Now().UnixNano())
	in := []string{"rock", "paper", "scissors"}
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	result = pick
	return result
}
*/
