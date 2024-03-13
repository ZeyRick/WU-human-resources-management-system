package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
)

type ReportController struct {
	service *services.ReportService
}

func NewReportController() *ReportController {
	return &ReportController{
		service: services.NewReportService(),
	}
}

func (ctrl *ReportController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.ReportFilter](r)
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
	https.ResponseJSON(w, r, http.StatusOK, result)
}

func (ctrl *ReportController) ExportList(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.ReportFilter](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	ctrl.service.Export(w, r, &pageOpt, &dto)
}
