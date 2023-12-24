package services

import (
	"backend/adapters/dtos"
	"backend/core/models/clock"
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
	err := srv.repo.Create(&clock.Clock{ EmployeeId: payload.EmployeeId, ClockType: payload.ClockType})
	return err
}
