package services

import (
	"backend/adapters/dtos"
	"backend/core/models/employee"
	"backend/core/models/schedule"
	"backend/core/types"
	"backend/pkg/logger"
	"encoding/json"
	"errors"
)

type ScheduleService struct {
	repo         *schedule.ScheduleRepo
	employeeRepo *employee.EmployeeRepo
}

func NewScheduleService() *ScheduleService {
	return &ScheduleService{
		repo:         schedule.NewScheduleRepo(),
		employeeRepo: employee.NewEmployeeRepo(),
	}
}

func (srv *ScheduleService) List(pageOpt *dtos.PageOpt, dto *dtos.ScheduleFilter) ( *types.ListData[schedule.Schedule] ,error) {
	result, err := srv.repo.List(pageOpt, dto)
	return result, err
}

func (srv *ScheduleService) Add(dto *dtos.AddSchedule) error {
	employee, err := srv.employeeRepo.FindId(dto.EmployeeId)
	if err != nil {
		return err
	}
	if employee.ID == 0 {
		return errors.New("user not exist")
	}
	datesJson, err := json.Marshal(dto.Dates)
	if err != nil {
		return err
	}
	existedSchedue, err := srv.repo.FindExistedScope(dto.EmployeeId, dto.Scope)
	if err != nil {
		return err
	}
	logger.Console(existedSchedue)
	if existedSchedue.ID != 0 {
		return errors.New("scope for this employee already exist")
	}
	err = srv.repo.Create(&schedule.Schedule{
		EmployeeId: dto.EmployeeId,
		Scope:      dto.Scope,
		Dates:      string(datesJson),
	})
	return err
}
