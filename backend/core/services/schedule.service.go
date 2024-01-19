package services

import (
	"backend/adapters/dtos"
	"backend/core/models/employee"
	"backend/core/models/schedule"
	"backend/core/types"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/times"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func (srv *ScheduleService) List(pageOpt *dtos.PageOpt, dto *dtos.ScheduleFilter) (*types.ListData[schedule.Schedule], error) {
	result, err := srv.repo.List(pageOpt, dto)
	return result, err
}

func (srv *ScheduleService) GetAll(w http.ResponseWriter, r *http.Request, dto *dtos.ScheduleFilter) (*[]types.ScheduleInfo, error) {
	schedulesData, err := srv.repo.GetAllByScope(dto)
	if err != nil {
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return nil, err
	}
	dayInMonth, err := times.DaysInMonth(dto.Scope)
	if err != nil {
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return nil, err
	}
	var result []types.ScheduleInfo
	for i := 0; i < dayInMonth; i++ {
		var scheduleInfo types.ScheduleInfo
		scheduleInfo.Scope = fmt.Sprintf(`%s-%d`, dto.Scope, i+1)
		for _, mySchedule := range *schedulesData {
			var scheduleDates []int
			err := json.Unmarshal([]byte(mySchedule.Dates), &scheduleDates)
			if err != nil {
				https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
				return nil, err
			}
			for _, date := range scheduleDates {
				if date == i+1 {
					var employeeData = types.FormatedEmployee{
						Name: mySchedule.Employee.Name,
						DepartmentId: mySchedule.Employee.DepartmentId,
						ProfilePic: mySchedule.Employee.ProfilePic,
						ClockInTime: mySchedule.ClockInTime.String(),
						ClockOutTime: mySchedule.ClockOutTime.String(),
					}
					
					scheduleInfo.Employees = append(scheduleInfo.Employees, employeeData)
					
					break
				}
			}
		}

		result = append(result, scheduleInfo)
	}

	return &result, err
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
