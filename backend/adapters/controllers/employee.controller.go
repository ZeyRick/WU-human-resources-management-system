package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/excel"
	"backend/pkg/file"
	"backend/pkg/helper"
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

func (ctrl *EmployeeController) Edit(w http.ResponseWriter, r *http.Request) {
	employeeId, err := https.GetParamsID(r, "employeeId")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if employeeId == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing employee id")
		return
	}
	dto, err := https.GetBody[dtos.AddEmployee](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Edit(w, r, employeeId, &dto)
}

func (ctrl *EmployeeController) Add(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.AddEmployee](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Add(w, r, &dto)
}

func (ctrl *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	employeeId, err := https.GetParamsID(r, "employeeId")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if employeeId == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing employee id")
		return
	}
	employee, err := ctrl.service.GetOneById(employeeId)
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

func (ctrl *EmployeeController) UploadFiles(w http.ResponseWriter, r *http.Request) {
	fileName, err := file.SaveFile(r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, fileName)
}

func (ctrl *EmployeeController) ImportEmployeeExcel(w http.ResponseWriter, r *http.Request) {
	var employee *dtos.AddEmployee
	filename, file, err := file.GetExcelFile(r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	employees, err := excel.ReadEmployeeExcel(file)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	for i := 0; i < len(employees); i++ {
		employee = &employees[i]
		ctrl.service.Add(w, r, employee)
	}
	https.ResponseMsg(w, r, http.StatusCreated, filename)
}
