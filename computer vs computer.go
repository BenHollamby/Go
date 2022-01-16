package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ComputerOne struct {
	wins   int
	losses int
	ties   int
}

type ComputerTwo struct {
	wins   int
	losses int
	ties   int
}

func main() {

	zerone := ComputerOne{1, 2, 3}
	zerone.wins = 0
	zerone.losses = 0
	zerone.ties = 0

	zerotwo := ComputerTwo{1, 2, 3}
	zerotwo.wins = 0
	zerone.losses = 0
	zerotwo.ties = 0

	rand.Seed(time.Now().UnixNano())
	const a string = "rock"
	const b string = "paper"
	const c string = "scissors"

	var x [3]string
	x[0] = a
	x[1] = b
	x[2] = c

	fmt.Println("Welcome to computer vs computer, rock paper scissors Golang Edition!")
	fmt.Println("100 million rounds battling it out and.....")

	for totalgames := 0; totalgames < 100000000; totalgames++ {
		random1 := rand.Intn(len(x))
		random2 := rand.Intn(len(x))
		cc1 := random1
		cc2 := random2

		if cc1 == cc2 {
			zerone.ties++
			zerotwo.ties++
		}

		if cc1 == 0 && cc2 == 2 || cc1 == 1 && cc2 == 0 || cc1 == 2 && cc2 == 1 {
			zerone.wins++
			zerotwo.losses++
		}

		if cc1 == 0 && cc2 == 1 || cc1 == 1 && cc2 == 2 || cc1 == 2 && cc2 == 0 {
			zerotwo.wins++
			zerone.losses++
		}
	}

	if zerone.wins == zerotwo.wins {
		fmt.Println("COMPUTERS TIED!")
	}
	if zerone.wins > zerotwo.wins {
		fmt.Println("COMPUTER ONE WINS!")
		fmt.Println("\tComputer One Wins: ", zerone.wins, "\n\tComputer One Losses: ", zerone.losses, "\n\tComputer One Tied: ", zerone.ties)
		fmt.Println("Verses")
		fmt.Println("\tComputer Two Wins: ", zerotwo.wins, "\n\tComputer Two Losses: ", zerotwo.losses, "\n\tComputer Two Tied: ", zerotwo.ties)
	}
	if zerone.wins < zerotwo.wins {
		fmt.Println("COMPUTER TWO WINS!")
		fmt.Println("\tComputer Two Wins: ", zerotwo.wins, "\n\tComputer Two Losses: ", zerotwo.losses, "\n\tComputer Two Tied: ", zerotwo.ties)
		fmt.Println("Verses")
		fmt.Println("\tComputer One Wins: ", zerone.wins, "\n\tComputer One Losses: ", zerone.losses, "\n\tComputer One Tied: ", zerone.ties)
	}

}
