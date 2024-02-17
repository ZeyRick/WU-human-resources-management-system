package employee

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/department"
	"backend/core/types"
	"backend/pkg/db"
)

type Employee struct {
	models.BaseModel
	Name         string                `gorm:"type:string;not null"`
	ProfilePic   string                `gorm:"type:string;not null"`
	TelegramID   int64                 `gorm:"type:int;not null"`
	DepartmentId *int                  `json:"departmentId" gorm:"type:number;not null"`
	Department   department.Department `json:"department"`
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

func (repo *EmployeeRepo) UpdateById(employee *Employee) (int64, error) {
	result := db.Database.Model(&Employee{}).Where("id = ?", employee.ID).Updates(*employee)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *EmployeeRepo) Create(newEmployee *Employee) error {
	result := db.Database.Create(newEmployee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) Delete(employeeId *int) error {
	result := db.Database.Delete(&Employee{}, *employeeId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) GetOneById(employeeId *int) (*Employee, error) {
	var data Employee
	result := db.Database.Limit(1).Find(&data, *employeeId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *EmployeeRepo) GetOneByName(name string) (*Employee, error) {
	var data Employee
	result := db.Database.Where("name = ?", name).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *EmployeeRepo) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeFilter) (*types.ListData[Employee], error) {
	query := db.Database.Joins(`JOIN departments ON employees.department_id = departments.id`).Preload("Department")
	if dto.DepartmentId != nil && *dto.DepartmentId != 0 {
		query = query.Where("employees.department_id = ?", *dto.DepartmentId)
	}
	if dto.EmployeeName != "" {
		query = query.Where(`name LIKE ?`, "%"+dto.EmployeeName+"%")
	}
	return models.List[Employee](pageOpt, query, "employees")
}

func (repo *EmployeeRepo) All(dto *dtos.EmployeeFilter) (*[]Employee, error) {
	var data []Employee
	query := db.Database
	if dto.DepartmentId != nil {
		query = query.Where("department_id = ?", *dto.DepartmentId)
	}
	if dto.EmployeeId != nil {
		query = query.Where("id = ?", *dto.EmployeeId)
	}
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}

func (repo *EmployeeRepo) FindTelegramId(telegramId *int64) (*Employee, error) {
	var data Employee
	result := db.Database.Where("telegram_id = ?", telegramId).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}
