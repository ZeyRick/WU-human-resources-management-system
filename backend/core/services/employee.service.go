package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type EmployeeService struct {
	repo *employee.EmployeeRepo
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		repo: employee.NewEmployeeRepo(),
	}
}

func (srv *EmployeeService) Edit(w http.ResponseWriter, r *http.Request, employeeId *int, payload *dtos.AddEmployee) {
	_, err := srv.repo.UpdateById(&employee.Employee{
		BaseModel:    models.BaseModel{ID: uint(*employeeId)},
		Name:         payload.Name,
		DepartmentId: &payload.DepartmentId,
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
	err := srv.repo.Create(&employee.Employee{
		Name:         payload.Name,
		ProfilePic:   payload.ProfilePic,
		DepartmentId: &payload.DepartmentId,
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

func (srv *EmployeeService) All(dto *dtos.EmployeeFilter) (*[]employee.Employee, error) {
	return srv.repo.All(dto)
}
