package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func main() {
	x := NewCar(2, 5)
	y := NewTrack(5)
	fmt.Printf("Speed: %v, batteryDrain: %v, battery: %v, distance: %v\n", x.speed, x.batteryDrain, x.battery, x.distance)
	fmt.Printf("Distance: %v", y.distance)

}

func NewCar(speed, batteryDrain int) Car {
	car := Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
	return car
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	track := Track{
		distance: distance,
	}
	return track
}
