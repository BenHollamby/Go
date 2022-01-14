package main

import (
	"fmt"
)

type Stargate struct {
	seasons    int
	actorName  string
	companions []string
}

func main() {
	aStargate := Stargate{
		seasons:   10,
		actorName: "Colonel O'Neil",
		companions: []string{
			"Teal'C",
			"Samantha Carter",
			"Dr Daniel Jackson",
		},
	}
	fmt.Println(aStargate)
	fmt.Println(aStargate.actorName)
	fmt.Println("Seasons: ", aStargate.seasons)
	fmt.Println("Companion 1:", aStargate.companions[0])
	fmt.Println("Companion 2:", aStargate.companions[1])
	fmt.Println("Companion 3:", aStargate.companions[2])
	fmt.Printf("%v was in %v seasons, and had %v, %v, and %v as companions.", aStargate.actorName, aStargate.seasons, aStargate.companions[0], aStargate.companions[1], aStargate.companions[2])
}
