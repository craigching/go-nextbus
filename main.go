package main

import (
	"fmt"
	"os"

	"github.com/cching/go-nextbus/services"
)

func main() {

	service := services.NewMetroTransitService()
	routes, err := service.GetRoutes()

	handleError(err)

	// fmt.Printf("%+v\n", routes)

	route := routes[0]
	directions, err := service.GetValidDirections(route)

	handleError(err)

	fmt.Println(directions)

	dir := directions[0]

	stops, err := service.GetStops(route, dir)

	handleError(err)

	fmt.Println(stops)

	stop := stops[0]

	departures, err := service.GetDepartures(route, dir, stop)

	handleError(err)

	fmt.Println(departures)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
