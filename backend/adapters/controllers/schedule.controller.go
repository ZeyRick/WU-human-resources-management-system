package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
	"time"
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
	for _, day := range body.Dates {
		if day < 1 || day > 31 {
			https.ResponseError(w, r, http.StatusInternalServerError, "Dates must be days betwenn 1-31")
			return
		}
	}
	err = ctr.service.Add(&body)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Somthing went wrong")
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Schedule created")
}

func (ctrl *ScheduleController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.ScheduleFilter](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	result, err := ctrl.service.List(&pageOpt, &dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}

func (ctrl *ScheduleController) GetAll(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.ScheduleFilter](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	result, err := ctrl.service.GetAll(w, r, &dto)
	if err != nil {
		logger.Trace(err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}
