package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/repos"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type CourseService struct {
	repo *repos.CourseRepo
}

func NewCourseService() *CourseService {
	return &CourseService{
		repo: repos.NewCourseRepo(),
	}
}

func (srv *CourseService) All(dto *dtos.CourseFilter) (*[]models.Course, error) {
	return srv.repo.All(dto)
}

func (srv *CourseService) GetByEmployee(employeeId uint) ([]models.Course, error) {
	return srv.repo.GetByEmployee(&employeeId)
}

func (srv *CourseService) List(pageOpt *dtos.PageOpt, dto *dtos.CourseFilter) (*types.ListData[models.Course], error) {
	return srv.repo.List(pageOpt, dto)
}

func (srv *CourseService) Add(w http.ResponseWriter, r *http.Request, payload *dtos.AddCourse) {
	err := srv.repo.Create(&models.Course{
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
	_, err := srv.repo.UpdateById(&models.Course{
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
