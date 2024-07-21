package hotelbeds

type HotelAvailabilityRequest struct {
	Stay        Stay        `json:"stay"`
	Occupancies []Occupancy `json:"occupancies"`
	Hotels      HotelsObj   `json:"hotels"`
}

// HotelsObj hold the weird struct defined in the api
// hence this weird name
type HotelsObj struct {
	Hotel []int `json:"hotel"`
}

type Stay struct {
	CheckIn  string `json:"checkIn"`
	CheckOut string `json:"checkOut"`
}

type Occupancy struct {
	Rooms    int `json:"rooms"`
	Adults   int `json:"adults"`
	Children int `json:"children"`
}
