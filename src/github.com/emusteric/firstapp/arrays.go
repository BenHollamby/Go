package main

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 85, 93, 4} //start with number in array. declare type. Can only contain one type
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Grades: %v\n", students)

	students[0] = "ben"
	fmt.Printf("Grades: %v\n", students)

	students[1] = "aaron"
	students[2] = "tim"
	fmt.Printf("Grades: %v\n", students)

	fmt.Printf("Student #1: %v\n", students[0])
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Student #3: %v\n", students[2])

	fmt.Printf("Number of Students: %v\n", len(students))
}
