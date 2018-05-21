package services

import "fmt"

const (
	routesURL           = "http://svc.metrotransit.org/NexTrip/Routes"
	directionsURLFormat = "http://svc.metrotransit.org/NexTrip/Directions/%s"
	stopsURLFormat      = "http://svc.metrotransit.org/NexTrip/Stops/%s/%d"
	departuresURLFormat = "http://svc.metrotransit.org/NexTrip/%s/%d/%s"
)

// GetRoutesURL returns the URL to be used to get all routes in
// NexTrip
func GetRoutesURL() string {
	return routesURL
}

// GetDirectionsURL returns the URL used to get the valid directions
// for a route in NexTrip
func GetDirectionsURL(route Route) string {
	return fmt.Sprintf(directionsURLFormat, route.Route)
}

// GetStopsURL returns the URL used to get the stops for the given
// route and direction
func GetStopsURL(route Route, dir Direction) string {
	return fmt.Sprintf(stopsURLFormat, route.Route, dir)
}

// GetDeparturesURL returns the URL used to get the departures for the
// given route, direction, and stop
func GetDeparturesURL(route Route, dir Direction, stop Stop) string {
	return fmt.Sprintf(departuresURLFormat, route.Route, dir, stop.Value)
}
