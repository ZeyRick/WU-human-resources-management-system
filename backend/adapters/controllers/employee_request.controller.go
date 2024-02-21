package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/telegrambot"
	"net/http"
)

type EmployeeRequestController struct {
	service *services.EmployeeRequestService
}

func NewEmployeeRequestController() *EmployeeRequestController {
	return &EmployeeRequestController{
		service: services.NewEmployeeRequestService(),
	}
}

func (ctrl *EmployeeRequestController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.EmployeeRequestFilter](r)
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

func (ctrl *EmployeeRequestController) Confirmation(w http.ResponseWriter, r *http.Request) {
	body, err := https.GetBody[dtos.Confirmation](r)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	ok, id, err := ctrl.service.Confirmation(body)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	if !ok {
		https.ResponseMsg(w, r, http.StatusOK, "Request Rejected")
		telegrambot.SendEmployeeRejectedMessage(id)
		return
	}
	https.ResponseMsg(w, r, http.StatusUnauthorized, "New Employee Added")
	telegrambot.SendEmployeeAddedMessage(id)
}
