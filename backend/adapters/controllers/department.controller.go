package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
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
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	result, err := ctrl.service.All(&dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
}
