package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/times"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
)

type ScheduleController struct {
	service *services.ScheduleService
}

func NewScheduleController() *ScheduleController {
	return &ScheduleController{
		service: services.NewScheduleService(),
	}
}

func (ctr *ScheduleController) Add(w http.ResponseWriter, r *http.Request) {
	body, err := https.GetBody[dtos.AddSchedule](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	_, err = time.Parse("2006-01", body.Scope)
	if err != nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Scope must be in time format of 'YYYY-MM'")
		return
	}
	clockInTime, err := times.ParseTime(body.ClockInTime)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	clockoutTime, err := times.ParseTime(body.ClockOutTime)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	dates := strings.Split(body.Dates, ",")
	for _, day := range dates {
		intDay, err := strconv.Atoi(day)
		if err != nil {
			https.ResponseError(w, r, http.StatusInternalServerError, "Invalid Dates: "+day)
			return
		}
		if intDay < 1 || intDay > 31 {
			https.ResponseError(w, r, http.StatusInternalServerError, "Dates must be days betwenn 1-31")
			return
		}
	}
	ctr.service.Add(w, r, &types.AddSchedule{
		EmployeeId:   body.EmployeeId,
		DepartmentId: body.DepartmentId,
		Scope:        body.Scope,
		Dates:        body.Dates,
		ClockInTime:  clockInTime,
		ClockOutTime: clockoutTime,
	})
}

func (ctr *ScheduleController) Update(w http.ResponseWriter, r *http.Request) {
	body, err := https.GetBody[dtos.AddSchedule](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	_, err = time.Parse("2006-01", body.Scope)
	if err != nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Scope must be in time format of 'YYYY-MM'")
		return
	}
	clockInTime, err := times.ParseTime(body.ClockInTime)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	clockoutTime, err := times.ParseTime(body.ClockOutTime)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	dates := strings.Split(body.Dates, ",")
	for _, day := range dates {
		intDay, err := strconv.Atoi(day)
		if err != nil {
			https.ResponseError(w, r, http.StatusInternalServerError, "Invalid Dates: "+day)
			return
		}
		if intDay < 1 || intDay > 31 {
			https.ResponseError(w, r, http.StatusInternalServerError, "Dates must be days betwenn 1-31")
			return
		}
	}
	ctr.service.Update(w, r, &types.UpdateSchedule{
		EmployeeId:   body.EmployeeId,
		DepartmentId: body.DepartmentId,
		Scope:        body.Scope,
		Dates:        body.Dates,
		ClockInTime:  clockInTime,
		ClockOutTime: clockoutTime,
	})
}

func (ctrl *ScheduleController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.ScheduleFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	result, err := ctrl.service.List(&pageOpt, &dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}

func (ctrl *ScheduleController) GetAllWithFormat(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.ScheduleFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	result, err := ctrl.service.GetAllWithFormat(w, r, &dto)
	if err != nil {
		logger.Trace(err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}

func (ctrl *ScheduleController) GetByEmployeeId(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.ScheduleFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	employeeId := chi.URLParam(r, "employeeId")
	if err != nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing employee ID")
		return
	}
	employeeIdInt, err := strconv.Atoi(employeeId)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	dto.EmployeeId = &employeeIdInt
	ctrl.service.GetByEmployeeId(w, r, &dto)
}
