package client

import "github.com/Gealber/nuitee/client/hotelbeds"

type NuiteeProvider interface {
	Availability(availability *hotelbeds.HotelAvailabilityRequest) (*hotelbeds.HotelAvailabilityResponse, error)
}
