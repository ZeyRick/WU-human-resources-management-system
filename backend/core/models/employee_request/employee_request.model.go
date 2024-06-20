package employee_request

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)

type EmployeeRequest struct {
	ID               uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       uint              `json:"employeeId" gorm:"type:string"`
	Employee         employee.Employee `json:"employee"`
	TelegramID       int64             `json:"telegramId" gorm:"type:int;not null"`
	TelegramUsername string            `json:"telegramUsername" gorm:"type:string"`
	CreatedAt        time.Time         `json:"createdAt"`
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

func (repo *EmployeeRequestRepo) FindId(Id *int) (EmployeeRequest, error) {
	var data EmployeeRequest
	result := db.Database.Limit(1).Find(&data, *Id)
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
	query := db.Database.Joins("JOIN employees ON employees.id = employee_requests.employee_id").Preload("Employee")
	if dto.EmployeeID != 0 {
		query = query.Where("employee_requests.employee_id = ?", dto.EmployeeID)
	}
	if dto.EmployeeName != "" {
		query = query.Where(`employees.name LIKE ?`, "%"+dto.EmployeeName+"%")
	}
	if dto.CourseId != 0 {
		query = query.Where("employees.course_id = ?", dto.EmployeeID)
	}
	if dto.TelegramUsername != "" {
		query = query.Where(`telegram_username LIKE ?`, "%"+dto.TelegramUsername+"%")
	}
	return models.List[EmployeeRequest](pageOpt, query, "employee_requests")
}

func (repo *EmployeeRequestRepo) Delete(Id *int) error {
	result := db.Database.Delete(&EmployeeRequest{}, *Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
