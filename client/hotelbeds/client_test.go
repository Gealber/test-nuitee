package hotelbeds

import (
	"fmt"
	"testing"

	"github.com/Gealber/nuitee/config"
)

func Test_Availability(t *testing.T) {
	cfg := config.Config()
	clt, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	hotelsID := []int{129410, 105360, 106101, 1762514, 106045, 1773908, 105389, 1790375, 1735444, 1780872, 1717734, 105406, 105328, 229436, 105329, 1753277}
	availability := HotelAvailabilityRequest{
		Stay: Stay{
			CheckIn:  "2024-08-25",
			CheckOut: "2024-08-26",
		},
		Occupancies: []Occupancy{
			{
				Rooms:  2,
				Adults: 2,
			},
			{
				Rooms:  1,
				Adults: 1,
			},
		},
		Hotels: HotelsObj{
			Hotel: hotelsID,
		},
	}
	resp, err := clt.Availability(&availability)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RESP: %+v\n", resp)
}

func Test_signature(t *testing.T) {
	cfg := config.Config()
	var ts int64 = 1721547522
	expectedSignature := "c7622b3a1323084aaab4a2d234e86c0b602e2391f763e0e343d03dac5f6ba316"
	got := signature(cfg.Hotelbeds.API, cfg.Hotelbeds.Secret, ts)

	if got != expectedSignature {
		t.Fatal(fmt.Sprintf("signaturee differs got: %s want: %s", got, expectedSignature))
	}
}
