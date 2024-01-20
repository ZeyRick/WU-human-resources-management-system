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
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	result, err := ctrl.service.All(&dto)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, http.StatusInternalServerError, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
}
