package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
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
	err = ctr.clockService.Clock(clockDto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
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
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	result, err := ctrl.clockService.List(&pageOpt, &dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}
