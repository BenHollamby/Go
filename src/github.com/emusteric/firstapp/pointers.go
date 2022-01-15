package main

import (
	"fmt"
)

func main() {
	a := 42
	b := a
	fmt.Println(a, b)

	var x int = 66
	var y *int = &x
	fmt.Println(x, y)

	var s int = 96
	var t *int = &s
	fmt.Println(s, *t)

	n := [3]int{1, 2, 3}
	m := &n[0]
	v := &n[1]
	fmt.Printf("%v %p %p\n", n, m, v)
}
