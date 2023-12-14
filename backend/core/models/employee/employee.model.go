package employee

import (
	"backend/core/models"
	"backend/pkg/db"
)

type Employee struct {
	models.BaseModel
	Name       string `gorm:"type:string;not null"`
	ProfilePic string `gorm:"type:string;not null"`
}

type EmployeeRepo struct{}

func NewEmployeeRepo() *Employee {
	return &Employee{}
}

func (repo *Employee) Create(newEmployee *Employee) error {
	result := db.Database.Create(newEmployee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *Employee) FindById(userId *uint) (Employee, error) {
	user := Employee{}
	result := db.Database.Limit(1).Find(&user, *userId)
	if result.Error != nil {
		return Employee{}, result.Error
	}
	return user, nil
}

func (repo *EmployeeRepo) DeleteById(userId *uint) (int64, error) {
	result := db.Database.Delete(&Employee{}, *userId)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *EmployeeRepo) UpdateById(user *Employee) (int64, error) {
	result := db.Database.Model(&Employee{}).Where("id = ?", user.ID).Updates(*user)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *EmployeeRepo) FindByEmployeeName(userName string) (Employee, error) {
	employee := Employee{}
	result := db.Database.Where("username = ?", userName).Limit(1).Find(&employee)
	if result.Error != nil {
		return Employee{}, result.Error
	}
	return employee, nil
}
