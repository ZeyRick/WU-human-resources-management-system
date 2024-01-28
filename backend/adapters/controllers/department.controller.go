package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
)

type DepartmentController struct {
	service *services.DepartmentService
}

func NewDepartmentController() *DepartmentController {
	return &DepartmentController{
		service: services.NewDepartmentService(),
	}
}

func (ctrl *DepartmentController) All(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.DepartmentFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r,  err)
		return
	}
	result, err := ctrl.service.All(&dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r,  err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
}

func (ctrl *DepartmentController) Add(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.AddDepartment](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Add(w, r, &dto)
}

func (ctrl *DepartmentController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.DepartmentFilter](r)
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

func (ctrl *DepartmentController) Edit(w http.ResponseWriter, r *http.Request) {
	departmentId, err := https.GetParamsID(r, "departmentId")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if departmentId == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing department id")
		return
	}
	dto, err := https.GetBody[dtos.AddDepartment](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Edit(w, r, departmentId, &dto)
}
