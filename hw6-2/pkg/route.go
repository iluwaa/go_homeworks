package tour

import (
	"fmt"
	"math/rand"
)

type Route struct {
	RoutePoints  []*RoutePoint
	Vehicles     []*Vehicle
	DoneVehicles []*Vehicle
}

type RoutePoint struct {
	Name           string
	Distance       int
	PassengerCount int
	Visited        bool
}

func NewRoute() *Route {
	return &Route{make([]*RoutePoint, 0), make([]*Vehicle, 0), make([]*Vehicle, 0)}
}

func (route *Route) AddRoute(name string, distance int) {
	if len(route.RoutePoints) == 0 {
		distance = 0
	}
	route.RoutePoints = append(route.RoutePoints, &RoutePoint{Name: name, Distance: distance, PassengerCount: rand.Intn(300-1) + 1})
}

func (route *Route) ShowVehicles() {
	for i, vehicle := range route.Vehicles {
		fmt.Printf("%d: %s\n", i+1, ReflectVehicle(vehicle))
	}
}

func (route *Route) ShowRoute() {
	fmt.Printf("Today's route is ")
	for i, vehicle := range route.RoutePoints {
		add := ""
		if i == 0 {
			add = fmt.Sprintf("(Start)-> %dkm ->", route.RoutePoints[i+1].Distance)
		} else if i == len(route.RoutePoints)-1 {
			add = fmt.Sprintf("(End)\n")
		} else {
			add = fmt.Sprintf("-> %dkm ->", route.RoutePoints[i+1].Distance)
		}

		fmt.Printf("%s%s", vehicle.Name, add)
	}
}

func (route *Route) AddVehicle(vehicle Vehicle) {
	route.Vehicles = append(route.Vehicles, &vehicle)
}
