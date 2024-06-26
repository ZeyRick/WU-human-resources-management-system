package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
	"time"
)

type ClockController struct {
	clockService *services.ClockService
}

func NewClockController() *ClockController {
	return &ClockController{
		clockService: services.NewClockService(),
	}
}

func (ctr *ClockController) Clock(w http.ResponseWriter, r *http.Request) {
	clockDto, err := https.GetBody[dtos.Clock](r)
	if err != nil {
		logger.Trace(err)
		return
	}
	err = ctr.clockService.Clock(w, r, clockDto)
	if err != nil {
		logger.Trace(err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Clock Created")
}

func (ctrl *ClockController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.ClockFilter](r)
	if dto.Date != "" {
		if _, err := time.Parse("2006-01-02 15:04:05", dto.Date); err != nil {
			logger.Trace(err)
			https.ResponseError(w, r, http.StatusBadRequest, "Bad time format")
			return
		}
	}
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	result, err := ctrl.clockService.List(&pageOpt, &dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}

func (ctrl *ClockController) Attendence(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.AttendenceFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	result, err := ctrl.clockService.Attendence(&pageOpt, &dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}

func (ctrl *ClockController) AttendenceExport(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.AttendenceFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	ctrl.clockService.AttendenceExport(w, r, &pageOpt, &dto)
}

func (ctrl *ClockController) Update(w http.ResponseWriter, r *http.Request) {
	payload, err := https.GetBody[dtos.UpdateClock](r)
	if err != nil {
		logger.Trace(err)
		return
	}
	clockId, err := https.GetParamsID(r, "id")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	ctrl.clockService.Update(w, r, clockId, &payload)
}

func (ctr *ClockController) ManualClock(w http.ResponseWriter, r *http.Request) {
	clockDto, err := https.GetBody[dtos.ManualClock](r)
	if err != nil {
		logger.Trace(err)
		return
	}
	err = ctr.clockService.ManualClock(w, r, clockDto)
	if err != nil {
		logger.Trace(err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Clock Created")
}

func (ctrl *ClockController) UpdateManual(w http.ResponseWriter, r *http.Request) {
	payload, err := https.GetBody[dtos.UpdateClock](r)
	if err != nil {
		logger.Trace(err)
		return
	}
	clockId, err := https.GetParamsID(r, "id")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	ctrl.clockService.UpdateManual(w, r, clockId, &payload)
}
