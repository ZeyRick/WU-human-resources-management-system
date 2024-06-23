package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
)

type EmployeeRequestRepo struct{}

func (repo *EmployeeRequestRepo) Create(newEmployeeRequest *models.EmployeeRequest) error {
	result := db.Database.Create(newEmployeeRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRequestRepo) FindbyTelegramId(telegramID *int64) (models.EmployeeRequest, error) {
	var data models.EmployeeRequest
	result := db.Database.Where("telegram_id = ?", telegramID).Limit(1).Find(&data)
	if result.Error != nil {
		return models.EmployeeRequest{}, result.Error
	}
	return data, nil
}

func (repo *EmployeeRequestRepo) FindId(Id *int) (models.EmployeeRequest, error) {
	var data models.EmployeeRequest
	result := db.Database.Limit(1).Find(&data, *Id)
	if result.Error != nil {
		return models.EmployeeRequest{}, result.Error
	}
	return data, nil
}

func (repo *EmployeeRequestRepo) UpdateByTelegramId(newEmployeeRequest *models.EmployeeRequest) error {
	result := db.Database.Model(&models.EmployeeRequest{}).Where("telegram_id = ?", newEmployeeRequest.TelegramID).Updates(*newEmployeeRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRequestRepo) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeRequestFilter) (*types.ListData[models.EmployeeRequest], error) {
	query := db.Database.Joins("JOIN employees ON employees.id = employee_requests.employee_id").Preload("Employee")
	if dto.EmployeeID != 0 {
		query = query.Where("employee_requests.employee_id = ?", dto.EmployeeID)
	}
	if dto.EmployeeName != "" {
		query = query.Where(`employees.name LIKE ?`, "%"+dto.EmployeeName+"%")
	}
	
	if dto.TelegramUsername != "" {
		query = query.Where(`telegram_username LIKE ?`, "%"+dto.TelegramUsername+"%")
	}
	return List[models.EmployeeRequest](pageOpt, query, "employee_requests")
}

func (repo *EmployeeRequestRepo) Delete(Id *int) error {
	result := db.Database.Delete(&models.EmployeeRequest{}, *Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
