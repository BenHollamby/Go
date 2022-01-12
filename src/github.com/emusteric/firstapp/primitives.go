package main

import (
	"fmt"
)

func main() {
	//boolean
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	x := 1 == 1
	y := 1 == 2
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T", y, y)

	//
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	//string
	s := "this is a string"
	q := []byte(s)
	fmt.Printf("%v, %T\n", q, q)
}