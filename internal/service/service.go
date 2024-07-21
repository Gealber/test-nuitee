package service

import "github.com/Gealber/nuitee/internal/service/model"

type RateProvider interface {
	Get(filters model.RateFilters) (*model.RateResponse, error)
}
