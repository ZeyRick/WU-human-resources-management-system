package services

import (
	"backend/adapters/dtos"
	"backend/core/models/course"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type CourseService struct {
	repo *course.CourseRepo
}

func NewCourseService() *CourseService {
	return &CourseService{
		repo: course.NewCourseRepo(),
	}
}

func (srv *CourseService) All(dto *dtos.CourseFilter) (*[]course.Course, error) {
	return srv.repo.All(dto)
}

func (srv *CourseService) List(pageOpt *dtos.PageOpt, dto *dtos.CourseFilter) (*types.ListData[course.Course], error) {
	return srv.repo.List(pageOpt, dto)
}

func (srv *CourseService) Add(w http.ResponseWriter, r *http.Request, payload *dtos.AddCourse) {
	err := srv.repo.Create(&course.Course{
		Alias: payload.Alias,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Course already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Course created")
}

func (srv *CourseService) Edit(w http.ResponseWriter, r *http.Request, courseId *int, payload *dtos.AddCourse) {
	_, err := srv.repo.UpdateById(&course.Course{
		ID:    uint(*courseId),
		Alias: payload.Alias,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Course alias already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Course updated")
}
