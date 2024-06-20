package services

import (
	"backend/adapters/dtos"
	"backend/core/models/employee"
	"backend/core/models/schedule"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/times"
	"encoding/json"
	"fmt"
	"math"
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

func (srv *ScheduleService) GetAllWithFormat(w http.ResponseWriter, r *http.Request, dto *dtos.ScheduleFilter) (*[]types.ScheduleInfo, error) {
	schedulesData, err := srv.repo.GetAllByScope(dto)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return nil, err
	}
	dayInMonth, err := times.DaysInMonth(dto.Scope)
	if err != nil {
		helper.UnexpectedError(w, r, err)
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
				helper.UnexpectedError(w, r, err)
				return nil, err
			}
			for _, date := range scheduleDates {
				if date == i+1 {
					var employeeData = types.FormatedEmployee{
						Name:         mySchedule.Employee.Name,
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

func (srv *ScheduleService) GetByEmployeeId(w http.ResponseWriter, r *http.Request, dto *dtos.ScheduleFilter) {
	schedulesData, err := srv.repo.GetAllByScope(dto)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if len(*schedulesData) < 1 {
		https.ResponseError(w, r, http.StatusOK, "User has no schedule")
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, (*schedulesData)[0])
}

func (srv *ScheduleService) Add(w http.ResponseWriter, r *http.Request, dto *types.AddSchedule) {
	// employees, err := srv.employeeRepo.All(&dtos.EmployeeFilter{EmployeeId: dto.EmployeeIds, CourseId: dto.CourseId})
	// if err != nil {
	// 	helper.UnexpectedError(w, r,  err)
	// 	return
	// }
	// if len(*employees) < 1 {
	// 	https.ResponseError(w, r, http.StatusInternalServerError, "No user found")
	// 	return
	// }
	datesJson, err := json.Marshal(dto.Dates)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	var newSchedules []schedule.Schedule
	for _, curEmployeeId := range *dto.EmployeeIds {
		// existedSchedue, err := srv.repo.FindExistedScope(&curEmployeeId, dto.Scope)
		// if err != nil {
		// 	helper.UnexpectedError(w, r, err)
		// 	return
		// }
		// if existedSchedue.ID != 0 {
		// 	https.ResponseError(w, r, http.StatusInternalServerError, fmt.Sprintf(`"Employee ID: %d already has a schedule"`, curEmployeeId))
		// 	return
		// }
		minuteWorkPerDay := int(math.Round(dto.ClockOutTime.Sub(*dto.ClockInTime).Minutes())) - *dto.MinuteBreakTime
		converetedDates := "[" + string(datesJson)[1:len(string(datesJson))-1] + "]"
		newSchedules = append(newSchedules, schedule.Schedule{
			EmployeeId:        curEmployeeId,
			Scope:             dto.Scope,
			Dates:             converetedDates,
			ClockInTime:       *dto.ClockInTime,
			ClockOutTime:      *dto.ClockOutTime,
			MinuteWorkPerDay:  &minuteWorkPerDay,
			MinuteBreakPerDay: dto.MinuteBreakTime,
		})
	}

	err = srv.repo.BatchCreate(&newSchedules)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Schedule created")
}

func (srv *ScheduleService) Update(w http.ResponseWriter, r *http.Request, dto *types.UpdateSchedule) {
	datesJson, err := json.Marshal(dto.Dates)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	converetedDates := "[" + string(datesJson)[1:len(string(datesJson))-1] + "]"
	err = srv.repo.Update(*dto.EmployeeIds, &schedule.Schedule{
		Dates:        converetedDates,
		Scope:        dto.Scope,
		ClockInTime:  *dto.ClockInTime,
		ClockOutTime: *dto.ClockOutTime,
	})
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Schedule updated")
}
