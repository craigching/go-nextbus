package services

import (
	"encoding/json"
	"strings"
)

// MetroTransitService provides a programmatic interface to the Metro
// Transit NexTrip REST API
type MetroTransitService struct {
	httpClient *HTTPClient
}

// Route represents a route object in the metro transit services
// NexTrip API
type Route struct {
	Description string
	ProviderID  string
	Route       string
}

// Direction represents the valid directions for stops.  Unknown is
// never returned by the NexTrip API, but is used to represent an
// unknown state
type Direction int

const (
	unknown Direction = iota
	south
	east
	west
	north
)

// StringToDirection converts a string to a Direction
func StringToDirection(str string) Direction {
	switch strings.ToLower(str) {
	case "south", "southbound":
		return south
	case "east", "eastbound":
		return east
	case "west", "westbound":
		return west
	case "north", "northbound":
		return north
	}

	return unknown
}

type direction struct {
	Text  string
	Value string
}

// Stop represents a stop on a route for a valid direction
type Stop struct {
	Text  string
	Value string
}

// Departure provides the details about when a bus will depart from a
// given stop
type Departure struct {
	Actual           bool
	BlockNumber      int
	DepartureText    string
	DepartureTime    string
	Description      string
	Gate             string
	Route            string
	RouteDirection   string
	Terminal         string
	VehicleHeading   float32
	VehicleLatitude  float32
	VehicleLongitude float32
}

// NewMetroTransitService makes a new MetroTransitService
func NewMetroTransitService() MetroTransitService {
	// TODO Seems hacky?  Should NewHTTPClient return a pointer
	// instead?
	httpClient := NewHTTPClient()
	return MetroTransitService{
		httpClient: &httpClient,
	}
}

// GetRoutes returns all current routes
func (mts *MetroTransitService) GetRoutes() ([]Route, error) {
	url := GetRoutesURL()

	str, err := mts.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	var routes []Route
	err = json.Unmarshal([]byte(str), &routes)

	if err != nil {
		// Return an empty route so caller can choose to ignore
		// errors.

		// TODO Maybe routes is empty at this point if there was an
		// error?
		return []Route{}, err
	}

	return routes, nil
}

// GetValidDirections returns the valid directions (north, south,
// etc.) for a given Route
func (mts *MetroTransitService) GetValidDirections(route Route) ([]Direction, error) {
	url := GetDirectionsURL(route)

	str, err := mts.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	var dirs []direction
	err = json.Unmarshal([]byte(str), &dirs)

	if err != nil {
		return []Direction{}, err
	}

	var directions []Direction
	for _, dir := range dirs {
		directions = append(directions, StringToDirection(dir.Text))
	}

	return directions, nil
}

// GetStops returns the stops on the given route for the specified
// direction
func (mts *MetroTransitService) GetStops(route Route, direction Direction) ([]Stop, error) {
	url := GetStopsURL(route, direction)

	str, err := mts.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	var stops []Stop
	err = json.Unmarshal([]byte(str), &stops)

	return stops, err
}

// GetDepartures returns the departures for the given route, direction
// and stop
func (mts *MetroTransitService) GetDepartures(route Route, direction Direction, stop Stop) ([]Departure, error) {
	url := GetDeparturesURL(route, direction, stop)

	str, err := mts.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	var departures []Departure
	err = json.Unmarshal([]byte(str), &departures)

	return departures, err
}
