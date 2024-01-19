package services

import (
	"backend/adapters/dtos"
	"backend/core/models/department"
)

type DepartmentService struct {
	repo *department.DepartmentRepo
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{
		repo: department.NewDepartmentRepo(),
	}
}


func (srv *DepartmentService) All(dto *dtos.DepartmentFilter) (*[]department.Department,error) {
	return  srv.repo.All(dto)
}
