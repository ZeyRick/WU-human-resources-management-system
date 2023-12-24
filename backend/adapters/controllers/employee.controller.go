package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		service: services.NewEmployeeService(),
	}
}

func (ctrl *EmployeeController) List(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.ListEmployee](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	result, err := ctrl.service.List(&dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseJSON(w, r, http.StatusCreated, result)
	return
}

func (ctrl *EmployeeController) Add(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.AddEmployee](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = ctrl.service.Add(&dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Employee Created")
	return
}
