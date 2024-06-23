package controllers

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
)

type ClockSettingController struct {
	services *services.ClockSettingService
}

func NewClockSettingController() *ClockSettingController {
	return &ClockSettingController{
		services: services.NewClockSettingService(),
	}
}

func (ctr *ClockSettingController) Get(w http.ResponseWriter, r *http.Request) {
	clockSetting, err := ctr.services.Get()
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w,r, http.StatusOK, clockSetting)
}

func (ctrl *ClockSettingController) Save(w http.ResponseWriter, r *http.Request) {
	payload, err := https.GetBody[dtos.CreateClockSetting](r)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	ctrl.services.Save(w, r, &models.ClockSetting{
		Coordinate: payload.Coordinate,
		AllowTime: payload.AllowTime,
		ClockRange: payload.ClockRange,
	})
}
