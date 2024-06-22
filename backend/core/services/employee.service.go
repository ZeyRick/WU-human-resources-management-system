package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/course"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type EmployeeService struct {
	repo       *employee.EmployeeRepo
	courseRepo *course.CourseRepo
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		repo:       employee.NewEmployeeRepo(),
		courseRepo: course.NewCourseRepo(),
	}
}

func (srv *EmployeeService) Edit(w http.ResponseWriter, r *http.Request, employeeId *int, payload *dtos.AddEmployee) {
	_, err := srv.repo.UpdateById(&employee.Employee{
		BaseModel:     models.BaseModel{ID: uint(*employeeId)},
		Name:          payload.Name,
		EmployeeType:  payload.EmployeeType,
		Salary:        payload.Salary,
		IdNumber:      payload.IdNumber,
		IdFileName:    payload.IdFileName,
		PhotoFileName: payload.PhotoFileName,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Employee name already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Employee updated")
}

func (srv *EmployeeService) Add(w http.ResponseWriter, r *http.Request, payload *dtos.AddEmployee) {
	courses, err := srv.courseRepo.FindByIds(payload.CourseId)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	err = srv.repo.Create(&employee.Employee{
		Name:          payload.Name,
		EmployeeType:  payload.EmployeeType,
		Salary:        payload.Salary,
		IdNumber:      payload.IdNumber,
		IdFileName:    payload.IdFileName,
		PhotoFileName: payload.PhotoFileName,
		Courses:       courses,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Employee already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Employee created")
}

func (srv *EmployeeService) Delete(employeeId *int) error {
	return srv.repo.Delete(employeeId)
}

func (srv *EmployeeService) GetOneById(employeeId *int) (*employee.Employee, error) {
	return srv.repo.GetOneById(employeeId)
}

func (srv *EmployeeService) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeFilter) (*types.ListData[employee.Employee], error) {
	return srv.repo.List(pageOpt, dto)
}

func (srv *EmployeeService) All(dto *dtos.EmployeeFilter) (*[]types.EmployeeWithSchedule, error) {
	return srv.repo.All(dto)
}
