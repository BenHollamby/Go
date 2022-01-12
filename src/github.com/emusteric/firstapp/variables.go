package main

import "fmt"

var ( //this way you can assign a bunch of variables before the function. Must be declared with type
	characterName    string = "Kallor"
	characterNemesis string = "Anomander"
	title            string = "Old King"
	appearances      int    = 3
)

func main() {
	var y int //really long way
	y = 99
	fmt.Println(y)
	fmt.Printf("%v, %T\n", y, y) //value and type \n for new line

	var i int = 42 //long way
	fmt.Println(i)
	fmt.Printf("%v, %T\n", i, i) //value and type \n for new line

	x := 66 //short way
	fmt.Println(x)

	z := "malazan"
	fmt.Println(z)
	fmt.Printf("%v, %T\n", z, z) //value and type

	fmt.Println("The", title, characterName, "arch enemy", characterNemesis, "battle", appearances, "times!")
	fmt.Printf("%v, %T\n", characterName, characterName)
	fmt.Printf("%v, %T", appearances, appearances)
}
