package employee_request

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)

type EmployeeRequest struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       uint      `gorm:"type:string"`
	TelegramID       int64     `gorm:"type:int;not null"`
	TelegramUsername string    `gorm:"type:string"`
	CreatedAt        time.Time `json:"createdAt"`
}

type EmployeeRequestRepo struct{}

func (repo *EmployeeRequestRepo) Create(newEmployeeRequest *EmployeeRequest) error {
	result := db.Database.Create(newEmployeeRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRequestRepo) FindbyTelegramId(telegramID *int64) (EmployeeRequest, error) {
	var data EmployeeRequest
	result := db.Database.Where("telegram_id = ?", telegramID).Limit(1).Find(&data)
	if result.Error != nil {
		return EmployeeRequest{}, result.Error
	}
	return data, nil
}

func (repo *EmployeeRequestRepo) UpdateByTelegramId(newEmployeeRequest *EmployeeRequest) error {
	result := db.Database.Model(&EmployeeRequest{}).Where("telegram_id = ?", newEmployeeRequest.TelegramID).Updates(*newEmployeeRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRequestRepo) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeRequestFilter) (*types.ListData[EmployeeRequest], error) {
	query := db.Database
	if dto.EmployeeID != 0 {
		query = query.Where("employee_requests.employee_id = ?", dto.EmployeeID)
	}
	if dto.TelegramUsername != "" {
		query = query.Where(`telegram_username LIKE ?`, "%"+dto.TelegramUsername+"%")
	}
	return models.List[EmployeeRequest](pageOpt, query, "employee_requests")
}

func (repo *EmployeeRequestRepo) Delete(telegramId *int64) error {
	result := db.Database.Model(&EmployeeRequest{}).Where("telegram_id = ?", telegramId).Delete(&EmployeeRequest{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
