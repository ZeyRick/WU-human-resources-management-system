package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/variable"
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
	if dto.PageOpt.Page == nil || *dto.PageOpt.Page == 0 {
		dto.PageOpt.Page = variable.Create[int64](1)
	}
	if dto.PageOpt.Size == nil || *dto.PageOpt.Size == 0 {
		dto.PageOpt.Size = variable.Create[int64](10)
	}
	result, err := ctrl.service.List(&dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
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
