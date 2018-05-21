package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoutesURL(t *testing.T) {

	// Method Under Test
	routesURL := GetRoutesURL()

	assert.Equal(
		t,
		"http://svc.metrotransit.org/NexTrip/Routes",
		routesURL,
		"The URL to get all routes wasn't an expected value.")
}

func TestGetDirectionsURL(t *testing.T) {
	routeText := "MyTestRoute"
	route := Route{"", "", routeText}
	expectedURL := fmt.Sprintf("http://svc.metrotransit.org/NexTrip/Directions/%s", routeText)

	// Method Under Test
	directionsURL := GetDirectionsURL(route)

	assert.Equal(
		t,
		expectedURL,
		directionsURL,
		"The URL to get valid directions for a route wasn't an expected value")
}

func TestGetStopURL(t *testing.T) {
	routeText := "MyTestRoute"
	directionText := "south"
	direction := StringToDirection(directionText)
	route := Route{"", "", routeText}
	expectedURL := fmt.Sprintf(
		"http://svc.metrotransit.org/NexTrip/Stops/%s/%d",
		routeText,
		direction)

	// Method Under Test
	stopsURL := GetStopsURL(route, direction)

	assert.Equal(
		t,
		expectedURL,
		stopsURL,
		"The URL to get the stops for a route wasn't an expected value")

}

func TestGetDeparturesURL(t *testing.T) {
	routeText := "MyTestRoute"
	directionText := "south"
	stopText := "MyTestStop"
	direction := StringToDirection(directionText)
	route := Route{"", "", routeText}
	stop := Stop{"", stopText}
	expectedURL := fmt.Sprintf(
		"http://svc.metrotransit.org/NexTrip/%s/%d/%s",
		routeText,
		direction,
		stopText)

	// Method Under Test
	departuresURL := GetDeparturesURL(route, direction, stop)

	assert.Equal(
		t,
		expectedURL,
		departuresURL,
		"The URL to get the departures wasn't an expected value")
}
