package hotelbeds

type HotelAvailabilityResponse struct {
	Hotels struct {
		Hotels   []Hotel `json:"hotels"`
		CheckIn  string  `json:"checkIn"`
		Total    int     `json:"total"`
		CheckOut string  `json:"checkOut"`
	} `json:"hotels"`
}

type Hotel struct {
	Code     int    `json:"code"`
	MinRate  string `json:"minRate"`
	MaxRate  string `json:"maxRate"`
	Currency string `json:"currency"`
}
