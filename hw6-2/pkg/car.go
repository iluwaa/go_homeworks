package tour

import (
	"fmt"
)

type Car struct {
	*TransportState
	TransportCharacteristics
}

func (car *Car) Stop() {
	if !car.TransportState.IsRunnig {
		return
	}
	car.TransportState.IsRunnig = false
	car.TransportState.CurrentSpeed = 0
}

func (car *Car) Start() int {
	car.TransportState.IsRunnig = true
	car.TransportState.CurrentSpeed = car.TransportCharacteristics.SpeedStep - car.TransportState.PassengerCount*2
	return car.CurrentSpeed
}
func (car *Car) MaxPassengers() int {
	return car.PassengerCapasity
}

func (car *Car) AvailableSeats() int {
	return car.PassengerCapasity - car.PassengerCount
}

func (car *Car) TakePassengers(count int) int {
	availableSeats := car.AvailableSeats()
	if count > availableSeats {
		car.PassengerCount = car.PassengerCapasity
		return count - availableSeats
	} else {
		car.PassengerCount += count
		return 0
	}
}
func (car *Car) DropPassengers(count int) {
	car.PassengerCount -= count
}

func (car *Car) ChangeSpeed() int {
	if car.TransportState.IsRunnig {
		maxSpeedWithPassangers := car.TransportCharacteristics.MaxSpeed - car.TransportState.PassengerCount*20
		speedStepWithPassangers := car.TransportCharacteristics.SpeedStep - car.TransportState.PassengerCount*2
		fmt.Printf(" (MaxSpeed: %d, CurrentSped: %d, Step: %d), PassengersCount: %d\n", maxSpeedWithPassangers, car.CurrentSpeed, speedStepWithPassangers, car.PassengerCount)

		if car.TransportState.CurrentSpeed < maxSpeedWithPassangers && maxSpeedWithPassangers > (car.TransportState.CurrentSpeed+speedStepWithPassangers) {
			car.TransportState.CurrentSpeed += speedStepWithPassangers
		} else {
			car.TransportState.CurrentSpeed += maxSpeedWithPassangers - car.TransportState.CurrentSpeed
		}
	}
	return car.CurrentSpeed
}

func NewCar() *Car {
	return &Car{}
}
