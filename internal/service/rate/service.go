package rate

import (
	"net/http"

	"github.com/Gealber/nuitee/client"
	"github.com/Gealber/nuitee/errors"
	"github.com/Gealber/nuitee/internal/service/model"
)

type service struct {
	provider client.NuiteeProvider
}

func New(provider client.NuiteeProvider) *service {
	return &service{provider: provider}
}

func (s *service) Get(filters model.RateFilters) (*model.RateResponse, error) {
	req, err := filters.ToAvailability()
	if err != nil {
		return nil, errors.NewErrService(http.StatusUnprocessableEntity, err.Error())
	}

	availability, err := s.provider.Availability(req)
	if err != nil {
		return nil, errors.NewErrService(http.StatusInternalServerError, err.Error())
	}

	rates, err := model.FromAvailabilityToRateResponse(availability)
	if err != nil {
		return nil, errors.NewErrService(http.StatusInternalServerError, err.Error())
	}

	return rates, nil
}
