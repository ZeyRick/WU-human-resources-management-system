package services

import (
	"backend/adapters/dtos"
	"backend/core/models/clock"
	"backend/pkg/logger"
)

type ClockService struct {
	repo clock.ClockRepo
}

func NewClockService() *ClockService {
	return &ClockService{
		repo: *clock.NewClockRepo(),
	}
}

func (srv *ClockService) Clock(payload dtos.Clock) string {
	err := srv.repo.Create(&clock.Clock{ EmployeeId: payload.EmployeeId, ClockType: payload.ClockType})
	if err != nil {
		logger.Trace(err)
	}
	return "1"
}
