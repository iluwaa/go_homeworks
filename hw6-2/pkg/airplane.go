package tour

import "fmt"

type Airplane struct {
	*TransportState
	TransportCharacteristics
}

func (airplane *Airplane) Stop() {
	if !airplane.TransportState.IsRunnig {
		return
	}
	airplane.TransportState.IsRunnig = false
	airplane.TransportState.CurrentSpeed = 0
}

func (airplane *Airplane) Start() int {
	airplane.TransportState.IsRunnig = true
	airplane.TransportState.CurrentSpeed = airplane.TransportCharacteristics.SpeedStep
	return airplane.CurrentSpeed

}

func (airplane *Airplane) MaxPassengers() int {
	return airplane.PassengerCapasity
}

func (airplane *Airplane) AvailableSeats() int {
	return airplane.PassengerCapasity - airplane.PassengerCount
}

func (airplane *Airplane) TakePassengers(count int) int {
	availableSeats := airplane.AvailableSeats()
	if count > availableSeats {
		airplane.PassengerCount = airplane.PassengerCapasity
		return count - availableSeats
	} else {
		airplane.PassengerCount += count
		return 0
	}
}
func (airplane *Airplane) DropPassengers(count int) {
	airplane.PassengerCount -= count
}

func (airplane *Airplane) ChangeSpeed() int {
	if airplane.TransportState.IsRunnig {
		fmt.Printf(" (MaxSpeed: %d, CurrentSped: %d, Step: %d), PassengersCount: %d\n", airplane.MaxSpeed, airplane.CurrentSpeed, airplane.SpeedStep, airplane.PassengerCount)

		if airplane.TransportState.CurrentSpeed < airplane.TransportCharacteristics.MaxSpeed && airplane.TransportCharacteristics.MaxSpeed > (airplane.TransportState.CurrentSpeed+airplane.TransportCharacteristics.SpeedStep) {
			airplane.TransportState.CurrentSpeed += airplane.TransportCharacteristics.SpeedStep
		} else {
			airplane.TransportState.CurrentSpeed += airplane.TransportCharacteristics.MaxSpeed - airplane.TransportState.CurrentSpeed
		}
	}
	return airplane.TransportState.CurrentSpeed
}

func NewAirplane() *Airplane {
	return &Airplane{}
}
