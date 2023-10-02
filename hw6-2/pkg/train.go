package tour

import (
	"fmt"
)

type Train struct {
	*TransportState
	TransportCharacteristics
}

func (train *Train) Stop() {
	if !train.TransportState.IsRunnig {
		return
	}
	train.TransportState.IsRunnig = false
	train.TransportState.CurrentSpeed = 0
}

func (train *Train) Start() int {
	train.TransportState.IsRunnig = true
	train.TransportState.CurrentSpeed = train.TransportCharacteristics.SpeedStep - int(float32(train.TransportState.PassengerCount)*0.03)
	return train.CurrentSpeed
}
func (train *Train) MaxPassengers() int {
	return train.PassengerCapasity
}

func (train *Train) AvailableSeats() int {
	return train.PassengerCapasity - train.PassengerCount
}

func (train *Train) TakePassengers(count int) int {
	availableSeats := train.PassengerCapasity - train.PassengerCount
	fmt.Println(availableSeats)
	if count > availableSeats {
		train.PassengerCount = train.PassengerCapasity
		return count - availableSeats
	} else {
		train.PassengerCount += count
		return 0
	}
}
func (train *Train) DropPassengers(count int) {
	train.PassengerCount -= count
}

func (train *Train) ChangeSpeed() int {
	if train.TransportState.IsRunnig {
		maxSpeedWithPassangers := train.TransportCharacteristics.MaxSpeed - int(float32(train.TransportState.PassengerCount)*0.1)
		speedStepWithPassangers := train.TransportCharacteristics.SpeedStep - int(float32(train.TransportState.PassengerCount)*0.03)
		fmt.Printf(" (MaxSpeed: %d, CurrentSped: %d, Step: %d), PassengersCount: %d\n", maxSpeedWithPassangers, train.CurrentSpeed, speedStepWithPassangers, train.PassengerCount)

		if train.TransportState.CurrentSpeed < maxSpeedWithPassangers && maxSpeedWithPassangers > (train.TransportState.CurrentSpeed+speedStepWithPassangers) {
			train.TransportState.CurrentSpeed += speedStepWithPassangers
		} else {
			train.TransportState.CurrentSpeed += maxSpeedWithPassangers - train.TransportState.CurrentSpeed
		}
	}
	return train.CurrentSpeed
}

func NewTrain() *Train {
	return &Train{}
}
