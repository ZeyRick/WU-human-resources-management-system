package services

import (
	"backend/adapters/dtos"
	"backend/core/models/employee"
	"backend/core/models/schedule"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/times"
	"backend/pkg/variable"
	"encoding/json"
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
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return nil, err
	}
	dayInMonth, err := times.DaysInMonth(dto.Scope)
	if err != nil {
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
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
				helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
				return nil, err
			}
			for _, date := range scheduleDates {
				if date == i+1 {
					var employeeData = types.FormatedEmployee{
						Name:         mySchedule.Employee.Name,
						DepartmentId: mySchedule.Employee.DepartmentId,
						ProfilePic:   mySchedule.Employee.ProfilePic,
						ClockInTime:  mySchedule.ClockInTime.String(),
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

func (srv *ScheduleService) Add(w http.ResponseWriter, r *http.Request, dto *types.AddSchedule) {
	employees, err := srv.employeeRepo.All(&dtos.EmployeeFilter{EmployeeId: dto.EmployeeId, DepartmentId: dto.DepartmentId})
	if err != nil {
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	if len(*employees) < 1 {
		https.ResponseError(w, r, http.StatusInternalServerError, "No user found")
		return
	}
	datesJson, err := json.Marshal(dto.Dates)
	if err != nil {
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	var newSchedules []schedule.Schedule
	for _, curEmployee := range *employees {
		existedSchedue, err := srv.repo.FindExistedScope(variable.Create[int](int(curEmployee.ID)), dto.Scope)
		if err != nil {
			helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
			return
		}
		if existedSchedue.ID != 0 {
			https.ResponseError(w, r, http.StatusInternalServerError, "Employee name: "+curEmployee.Name+" already has a schedule")
			return
		}
		converetedDates := "[" + string(datesJson)[1:len(string(datesJson))-1] + "]"
		newSchedules = append(newSchedules, schedule.Schedule{
			EmployeeId:   variable.Create[int](int(curEmployee.ID)),
			Scope:        dto.Scope,
			Dates:        converetedDates,
			ClockInTime:  *dto.ClockInTime,
			ClockOutTime: *dto.ClockOutTime,
		})
	}

	err = srv.repo.BatchCreate(&newSchedules)
	if err != nil {
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Schedule created")
}
