package services

import (
	"backend/adapters/dtos"
	"backend/core/models/department"
	"backend/core/types"
	"backend/pkg/helper"
	"backend/pkg/https"
	"net/http"
	"strings"
)

type DepartmentService struct {
	repo *department.DepartmentRepo
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{
		repo: department.NewDepartmentRepo(),
	}
}

func (srv *DepartmentService) All(dto *dtos.DepartmentFilter) (*[]department.Department, error) {
	return srv.repo.All(dto)
}

func (srv *DepartmentService) List(pageOpt *dtos.PageOpt, dto *dtos.DepartmentFilter) (*types.ListData[department.Department], error) {
	return srv.repo.List(pageOpt, dto)
}

func (srv *DepartmentService) Add(w http.ResponseWriter, r *http.Request, payload *dtos.AddDepartment) {
	err := srv.repo.Create(&department.Department{
		Alias: payload.Alias,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Department already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Department created")
}

func (srv *DepartmentService) Edit(w http.ResponseWriter, r *http.Request, departmentId *int, payload *dtos.AddDepartment) {
	_, err := srv.repo.UpdateById(&department.Department{
		ID:    uint(*departmentId),
		Alias: payload.Alias,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			https.ResponseError(w, r, http.StatusBadRequest, "Department alias already existed")
			return
		}
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Department updated")
}
