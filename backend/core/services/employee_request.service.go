package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/models/employee_request"
	"backend/core/types"
	"backend/pkg/logger"
)

type EmployeeRequestService struct {
	emp  *employee.EmployeeRepo
	repo *employee_request.EmployeeRequestRepo
}

func NewEmployeeRequestService() *EmployeeRequestService {
	return &EmployeeRequestService{
		emp: employee.NewEmployeeRepo(),
	}
}

func Confirmation() {

}

func (srv *EmployeeRequestService) Pend(name string, id *int64, telegramName string) (bool, error) {
	employees, err := srv.emp.GetOneByName(name)
	if err != nil {
		logger.Trace(err)
		return false, err
	}
	if employees.ID == 0 {
		return false, nil
	}
	err = srv.repo.Create(&employee_request.EmployeeRequest{
		EmployeeID:       employees.ID,
		TelegramID:       int64(*id),
		TelegramUsername: telegramName,
	})
	if err != nil {
		logger.Trace(err)
		return false, err
	}
	return true, nil
}

func (srv *EmployeeRequestService) CheckPending(id *int64) (bool, error) {
	result, err := srv.repo.FindbyTelegramId(id)
	if err != nil {
		logger.Trace(err)
		return false, err
	}
	if result.TelegramID != 0 {
		return false, nil
	}
	return true, nil
}

func (srv *EmployeeRequestService) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeRequestFilter) (*types.ListData[employee_request.EmployeeRequest], error) {
	return srv.repo.List(pageOpt, dto)
}

func (srv *EmployeeRequestService) Confirmation(dto dtos.Confirmation) error {
	if dto.Confirmation == types.Rejected {
		return srv.repo.Delete(dto.TelegramID)
	}
	request, err := srv.repo.FindbyTelegramId(dto.TelegramID)
	if err != nil {
		return err
	}
	_, err = srv.emp.UpdateById(&employee.Employee{
		BaseModel:  models.BaseModel{ID: request.EmployeeID},
		TelegramID: *dto.TelegramID,
	})
	if err != nil {
		return err
	}
	return srv.repo.Delete(dto.TelegramID)
}

func (srv *EmployeeRequestService) Test() string {
	return "Hello"
}
