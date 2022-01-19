package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	GetItem(a, index)
}

func GetItem(slice []int, index int) {
	fmt.Println(slice[2])
	fmt.Println(len(int(index)))
}
