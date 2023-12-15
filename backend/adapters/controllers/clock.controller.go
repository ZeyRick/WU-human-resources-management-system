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

func (ctr *ClockController)Clock(w http.ResponseWriter, r *http.Request) {
	clockDto, err := https.GetBody[dtos.Clock](r);
	if err != nil {
		logger.Trace(err)
		return
	}
	status := ctr.clockService.Clock(clockDto)
	if ( status == "1") {
		https.ResponseJSON(w, r, http.StatusOK,"Hello World")
		return
	}
	https.ResponseText(w, r, http.StatusBadRequest,"No World")
	return
}
