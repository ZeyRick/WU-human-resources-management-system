package employee

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
)

type Employee struct {
	models.BaseModel
	Name       string `gorm:"type:string;not null"`
	ProfilePic string `gorm:"type:string;not null"`
}

type EmployeeRepo struct{}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{}
}

func (repo *EmployeeRepo) Create(newEmployee *Employee) error {
	result := db.Database.Create(newEmployee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) List(dto *dtos.ListEmployee)(*types.ListData[Employee], error) {
	return models.List[Employee](&dto.PageOpt, db.Database, "employees")
}