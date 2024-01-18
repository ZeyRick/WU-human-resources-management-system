package services

import (
	"backend/adapters/dtos"
	"backend/core/models/clock"
	"backend/core/types"
)

type ClockService struct {
	repo clock.ClockRepo
}

func NewClockService() *ClockService {
	return &ClockService{
		repo: *clock.NewClockRepo(),
	}
}

func (srv *ClockService) Clock(payload dtos.Clock) error {
	err := srv.repo.Create(&clock.Clock{EmployeeId: payload.EmployeeId, ClockType: payload.ClockType})
	return err
}

func (srv *ClockService) List(pageOpt *dtos.PageOpt, dto *dtos.ClockFilter) (*types.ListData[clock.Clock], error) {
	result, err := srv.repo.List(pageOpt, dto)
	return result, err
}
