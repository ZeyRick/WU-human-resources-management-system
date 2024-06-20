package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/logger"
	"net/http"
)

type CourseController struct {
	service *services.CourseService
}

func NewCourseController() *CourseController {
	return &CourseController{
		service: services.NewCourseService(),
	}
}

func (ctrl *CourseController) All(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.CourseFilter](r)
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

func (ctrl *CourseController) Add(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.AddCourse](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Add(w, r, &dto)
}

func (ctrl *CourseController) List(w http.ResponseWriter, r *http.Request) {
	pageOpt, dto, err := https.GetPaginationWithType[dtos.CourseFilter](r)
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

func (ctrl *CourseController) Edit(w http.ResponseWriter, r *http.Request) {
	courseId, err := https.GetParamsID(r, "courseId")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if courseId == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing course id")
		return
	}
	dto, err := https.GetBody[dtos.AddCourse](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.service.Edit(w, r, courseId, &dto)
}
