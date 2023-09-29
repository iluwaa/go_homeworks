package tour

import (
	"fmt"
	"math/rand"
	"os"
)

func Tour() {
	route := DummyRoute()
	route.ShowRoute()
	fmt.Println("We have a next vehicles on route:")
	route.ShowVehicles()
	fmt.Println()

	vehicle := route.Vehicles[rand.Intn(len(route.Vehicles))]

	route.Ride(*vehicle)

}

func (route *Route) Ride(vehicle Vehicle) {
	for i := 0; i < len(route.RoutePoints); i++ {
		var isLastRide bool

		if i+2 == len(route.RoutePoints) {
			isLastRide = true
		} else {
			isLastRide = false
		}

		fmt.Printf("In %s we have a %d passengers. ", route.RoutePoints[i].Name, route.RoutePoints[i].PassengerCount)

		leftover := vehicle.TakePassengers(route.RoutePoints[i].PassengerCount)

		if leftover != 0 {
			fmt.Printf("%s can take only %d passengers, so %d passengers will wait next vehicle.\n", ReflectVehicle(&vehicle), vehicle.MaxPassengers(), leftover)
		} else {
			fmt.Printf("All %d taken.\n", route.RoutePoints[i].PassengerCount)
		}
		fmt.Println()

		progress := vehicle.Start()
		time := 1

		for progress < route.RoutePoints[i+1].Distance {
			time++
			fmt.Printf("For %d hours %s passed %d kms", time, ReflectVehicle(&vehicle), progress)
			progress += vehicle.ChangeSpeed()

		}
		fmt.Printf("Path was done for %d hours\n", time+1)
		vehicle.Stop()

		fmt.Printf("\n----------\n")
		if isLastRide {
			PassengerToDrop := vehicle.MaxPassengers() - vehicle.AvailableSeats()
			vehicle.DropPassengers(PassengerToDrop)
			fmt.Printf("In %s we dropped %d passengers.\n", route.RoutePoints[i+1].Name, PassengerToDrop)
			route.DoneVehicles = append(route.DoneVehicles, &vehicle)
			fmt.Println("Route done!")
			os.Exit(0)
		} else {
			PassengerToDrop := rand.Intn(vehicle.MaxPassengers() - vehicle.AvailableSeats())
			vehicle.DropPassengers(PassengerToDrop)
			fmt.Printf("In %s we dropped %d passengers.\n", route.RoutePoints[i+1].Name, PassengerToDrop)
		}
	}
}

func DummyAirplane() *Airplane {
	airplane := NewAirplane()
	airplane.TransportCharacteristics = TransportCharacteristics{
		SpeedStep:         rand.Intn(150-100) + 100,
		MaxSpeed:          rand.Intn(700-500) + 500,
		PassengerCapasity: 100,
	}
	airplane.TransportState = &TransportState{}

	return airplane
}

func DummyCar() *Car {
	car := NewCar()
	car.TransportCharacteristics = TransportCharacteristics{
		SpeedStep:         rand.Intn(25-10) + 10,
		MaxSpeed:          rand.Intn(280-120) + 120,
		PassengerCapasity: 4,
	}
	car.TransportState = &TransportState{}

	return car
}

func DummyTrain() *Train {
	train := NewTrain()
	train.TransportCharacteristics = TransportCharacteristics{
		SpeedStep:         rand.Intn(100-90) + 90,
		MaxSpeed:          rand.Intn(350-300) + 300,
		PassengerCapasity: 500,
	}
	train.TransportState = &TransportState{}

	return train
}

func DummyRoute() *Route {
	route := NewRoute()

	route.AddRoute("Kyiv", 999999)
	route.AddRoute("Kharkiv", 489)
	route.AddRoute("Lviv", 1020)

	route.AddVehicle(DummyAirplane())
	route.AddVehicle(DummyCar())
	route.AddVehicle(DummyTrain())
	return route
}
