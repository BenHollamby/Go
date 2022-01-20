package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	GetItem(a)
	getLength(a)
	getLength(a[3:])
}

func GetItem(slice []int) {
	fmt.Println(slice[2])
}

func getLength(length []int) {
	fmt.Println(len(length))
}

/*
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	s := []string{"James", "Wagner", "Christene", "Mike"}
	fmt.Println(contains(s, "James")) // true
	fmt.Println(contains(s, "Jack")) // false
}
*/
