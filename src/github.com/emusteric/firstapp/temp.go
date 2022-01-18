package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	fmt.Println(strings.Repeat("*", numStarsPerLine))
	fmt.Println(welcomeMsg)
	fmt.Println(strings.Repeat("*", numStarsPerLine))
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	i := strings.ReplaceAll(oldMsg, "*", "")
	x := strings.TrimSpace(i)
	return x
	panic("Please implement the CleanupMessage() function")
}
