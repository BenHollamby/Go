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
	fmt.Printf("Grades: %v\n", students) //all students in an array

	fmt.Printf("Student #1: %v\n", students[0])
	fmt.Printf("Student #2: %v\n", students[1])
	fmt.Printf("Student #3: %v\n", students[2])

	fmt.Printf("Number of Students: %v\n", len(students)) //total number of student

	//loop array with index against length
	test := [6]string{"Stargate", "SG-1", "Atlantis", "Universe", "Voyager", "Next Generation"}
	for i := 0; i < len(test); i++ {
		fmt.Println(i, test[i])
	}

	//loop array with for Range
	for index, value := range test {
		fmt.Println(index, value)
	}

	//blank identifier, ignoring index
	for _, value := range test {
		fmt.Println(value)
	}

	//blank identifier, ignoring value
	for index, _ := range test {
		fmt.Println(index)
	}
}
