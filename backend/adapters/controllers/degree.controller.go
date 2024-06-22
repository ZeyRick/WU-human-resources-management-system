package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
)

type DegreeController struct {
	service *services.DegreeService
}

func NewDegreeControler() *DegreeController {
	return &DegreeController{
		service: services.NewDegreeService(),
	}

}
func (ctrl *DegreeController) All(w http.ResponseWriter, r *http.Request) {
	result, err := ctrl.service.All()
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
}

func (ctrl *DegreeController) Add(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.AddDegree](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Add(w, r, &dto)
}

func (ctrl *DegreeController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.DegreeFilter](r)
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

func (ctrl *DegreeController) Edit(w http.ResponseWriter, r *http.Request) {

	dto, err := https.GetBody[dtos.AddDegree](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Edit(w, r, &dto)
}
func (ctrl *DegreeController) SearchList(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetParamsStrV2(r, "alias")
	data := dtos.DegreeFilter{
		Alias: dto,
	}
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	result, err := ctrl.service.SearchList(&data)
	if err != nil {
		logger.Trace(err)
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, result)
}
