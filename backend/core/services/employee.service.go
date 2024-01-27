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
	return srv.repo.Create(&employee.Employee{
		Name:       payload.Name,
		ProfilePic: payload.ProfilePic,
	})
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

func (srv *EmployeeService) PendingList(pageOpt *dtos.PageOpt, dto *dtos.EmployeeFilter) (*types.ListData[employee.Employee], error) {
	return srv.repo.PendingList(pageOpt, dto)
}
