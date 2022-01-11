package main

import "fmt"

func main(){

var array[3]string

array[0] = "rock"
array[1] = "scissors"
array[2] = "paper"

fmt.Println("Items in my array are:")
fmt.Println("Item 1:", array[0])
fmt.Println("Item 2:", array[1])
fmt.Println("Item 3:", array[2])

}