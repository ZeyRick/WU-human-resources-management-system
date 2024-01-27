package services

import (
	"backend/core/models/employee"
	"backend/core/types"
	"fmt"
)

type ConfirmationService struct {
	repo *employee.EmployeeRepo
}

func NewConfirmationService() *ConfirmationService {
	return &ConfirmationService{
		repo: employee.NewEmployeeRepo(),
	}
}

func Confirmation() {

}

func (srv *ConfirmationService) Pend(id int64) error {
	fmt.Println("Hello")
	return srv.repo.Create(&employee.Employee{
		TelegramID: id,
		Status:     types.Pending,
	})
}

func (srv *ConfirmationService) PendingList(id int64) error {
	return srv.repo.Create(&employee.Employee{
		TelegramID: id,
		Status:     types.Pending,
	})
}
