package elon

import (
	"fmt"
	"math"
)

// Drive - define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery > car.batteryDrain-1 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance - define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

// DisplayBattery - define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", car.battery)
}

// CanFinish - define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	var numOfDrives = int(math.Ceil(float64(trackDistance) / float64(car.speed)))
	if car.battery > numOfDrives*car.batteryDrain-1 {
		return true
	} else {
		return false
	}
}
