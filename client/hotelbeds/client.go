package hotelbeds

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Gealber/nuitee/config"
)

const (
	// APIURL = "https://developer.hotelbeds.com"
	APIURL = "https://api.test.hotelbeds.com/hotel-api/1.0"
)

type client struct {
	api    string
	secret string
}

func New(cfg *config.AppConfig) (*client, error) {
	if cfg.Hotelbeds.API == "" || cfg.Hotelbeds.Secret == "" {
		return nil, errors.New("api or secret are not set on evironment variables, please check HOTELBEDS_SECRET and HOTELBEDS_API_KEY")
	}

	return &client{
		api:    cfg.Hotelbeds.API,
		secret: cfg.Hotelbeds.Secret,
	}, nil
}

func (c *client) Availability(availability *HotelAvailabilityRequest) (*HotelAvailabilityResponse, error) {
	data, err := json.Marshal(&availability)
	if err != nil {
		return nil, err
	}

	path := "hotels"

	var availabilityResp HotelAvailabilityResponse
	err = c.post(path, data, &availabilityResp)
	if err != nil {
		return nil, err
	}

	return &availabilityResp, nil
}

func (c *client) post(path string, data []byte, objResp any) error {
	clt := http.Client{}

	url := APIURL + "/" + path
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	signature := signature(c.api, c.secret, time.Now().UTC().Unix())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-key", c.api)
	req.Header.Set("X-Signature", signature)

	resp, err := clt.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(fmt.Sprintf("unexpected status code received: %d %s", resp.StatusCode, string(b)))
	}

	dataResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(dataResp, objResp)
}

func signature(apiKey, secret string, ts int64) string {
	data := fmt.Sprintf("%s%s%d", apiKey, secret, ts)
	sum := sha256.Sum256([]byte(data))

	return hex.EncodeToString(sum[:])
}
