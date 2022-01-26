package main

func main() {
	InterestRate(200.75)
}

func InterestRate(balance float64) float32 {
	var x float32
	if balance < 0 {
		x = 3.213
	} else if balance < 1000 {
		x = 0.5
	} else if balance > 1000 && balance < 5000 {
		x = 1.621
	} else if balance > 5000 {
		x = 2.475
	}
	return x
}
