package tour

import (
	"reflect"
	"strings"
)

type TransportCharacteristics struct {
	MaxSpeed          int
	SpeedStep         int
	PassengerCapasity int
}

type TransportState struct {
	IsRunnig       bool
	CurrentSpeed   int
	PassengerCount int
}

type Transport interface {
	Start() int
	Stop()
	ChangeSpeed() int
}

type Vehicle interface {
	TakePassengers(int) int
	DropPassengers(int)
	MaxPassengers() int
	AvailableSeats() int
	Transport
}

func ReflectVehicle(vehicle *Vehicle) string {
	return strings.Split(reflect.TypeOf(*vehicle).Elem().String(), ".")[1]
}
