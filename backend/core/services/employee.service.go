package services

import (
	"backend/adapters/dtos"
	"backend/core/models/employee"
	"backend/core/types"
)

type EmployeeService struct {
	repo *employee.EmployeeRepo
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		repo: employee.NewEmployeeRepo(),
	}
}


func (srv *EmployeeService) Add(payload *dtos.AddEmployee) error {
	err := srv.repo.Create(&employee.Employee{
		Name: payload.Name,
		ProfilePic: payload.ProfilePic,
	})
	return err
}

func (srv *EmployeeService) List(params *dtos.ListEmployee) ( *types.ListData[employee.Employee] ,error) {
	result, err := srv.repo.List(params)
	return result, err
}
