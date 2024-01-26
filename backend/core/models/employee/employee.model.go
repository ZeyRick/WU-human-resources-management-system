package employee

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
)

type Employee struct {
	models.BaseModel
	Name       string         `gorm:"type:string;not null"`
	ProfilePic string         `gorm:"type:string;not null"`
	TelegramID int64          `gorm:"type:int;not null"`
	Status     types.PendType `gorm:"type:ENUM;not null"`
}

type EmployeeRepo struct{}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{}
}

func (repo *EmployeeRepo) FindId(employeeId *int) (Employee, error) {
	var data Employee
	result := db.Database.Limit(1).Find(&data, *employeeId)
	if result.Error != nil {
		return Employee{}, result.Error
	}
	return data, nil
}

func (repo *EmployeeRepo) Create(newEmployee *Employee) error {
	result := db.Database.Create(newEmployee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeFilter) (*types.ListData[Employee], error) {
	return models.List[Employee](pageOpt, db.Database, "employees")
}

func (repo *EmployeeRepo) All(dto *dtos.EmployeeFilter) (*[]Employee, error) {
	var data []Employee
	query := db.Database
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}
