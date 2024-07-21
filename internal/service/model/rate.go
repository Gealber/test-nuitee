package model

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Gealber/nuitee/client/hotelbeds"
)

type RateResponse struct {
	Data     []HotelData `json:"data"`
	Supplier Suplier     `json:"supplier"`
}

type Suplier struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type HotelData struct {
	HotelID  string  `json:"hotelId"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}

type RateFilters struct {
	Checkin          string `form:"checkin"`
	Checkout         string `form:"checkout"`
	Currency         string `form:"currency"`
	GuestNationality string `form:"guestNationality"`
	HotelsIDs        string `form:"hotelIds"`
	Occupancies      string `form:"occupancies"`
}

type Occupancy struct {
	Rooms    int `form:"rooms"`
	Adults   int `form:"adults"`
	Children int `form:"children"`
}

func (filter RateFilters) ToAvailability() (*hotelbeds.HotelAvailabilityRequest, error) {
	splt := strings.Split(filter.HotelsIDs, ",")
	ids := make([]int, len(splt))
	for i, n := range splt {
		id, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}

	occupancies, err := ToHotelbedsOcupancies(filter.Occupancies)
	if err != nil {
		return nil, err
	}

	return &hotelbeds.HotelAvailabilityRequest{
		Stay: hotelbeds.Stay{
			CheckIn:  filter.Checkin,
			CheckOut: filter.Checkout,
		},
		Occupancies: occupancies,
		Hotels: hotelbeds.HotelsObj{
			Hotel: ids,
		},
	}, nil
}

func ToHotelbedsOcupancies(data string) ([]hotelbeds.Occupancy, error) {
	var occupancies []hotelbeds.Occupancy
	err := json.Unmarshal([]byte(data), &occupancies)
	if err != nil {
		return nil, err
	}

	result := make([]hotelbeds.Occupancy, len(occupancies))
	for i := range occupancies {
		result[i] = hotelbeds.Occupancy{
			Rooms:    occupancies[i].Rooms,
			Adults:   occupancies[i].Adults,
			Children: occupancies[i].Children,
		}
	}

	return result, nil
}

func FromAvailabilityToRateResponse(availability *hotelbeds.HotelAvailabilityResponse) (*RateResponse, error) {
	data := make([]HotelData, len(availability.Hotels.Hotels))
	hotels := availability.Hotels.Hotels
	for i := range hotels {
		price, err := strconv.ParseFloat(hotels[i].MaxRate, 64)
		if err != nil {
			return nil, err
		}

		data[i] = HotelData{
			HotelID:  strconv.Itoa(hotels[i].Code),
			Currency: hotels[i].Currency,
			Price:    price,
		}
	}

	return &RateResponse{
		Data: data,
	}, nil
}
