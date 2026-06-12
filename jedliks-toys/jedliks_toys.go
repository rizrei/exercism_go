package jedlik

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

func (car *Car) Drive() {
	if car.canDrive() {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car Car) CanFinish(trackDistance int) bool {
	car.distance = 0
	for car.distance < trackDistance {
		if car.canDrive() {
			car.Drive()
		} else {
			return false
		}
	}
	return true
}

func (car Car) canDrive() bool {
	return car.battery >= car.batteryDrain
}
