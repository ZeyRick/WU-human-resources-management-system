package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/variable"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type EmployeeController struct {
	service *services.EmployeeService
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		service: services.NewEmployeeService(),
	}
}

func (ctrl *EmployeeController) All(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.EmployeeFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	result, err := ctrl.service.All(&dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
}

func (ctrl *EmployeeController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.EmployeeFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	result, err := ctrl.service.List(&pageOpt, &dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
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
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Employee Created")
}

func (ctrl *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	employeeIdStr := chi.URLParam(r, "employeeId")
	if employeeIdStr == "" {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing employee id")
		return
	}
	employeeId, err := strconv.Atoi(employeeIdStr)
	if err != nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Invalid employee id")
		return
	}
	employee, err := ctrl.service.GetOneById(&employeeId)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if employee == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Employee not found")
		return
	}
	err = ctrl.service.Delete(variable.Create[int](int(employee.ID)))
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Employee deleted")
}
