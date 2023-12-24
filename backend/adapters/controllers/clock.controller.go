package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
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
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Clock Created")
	return
}
