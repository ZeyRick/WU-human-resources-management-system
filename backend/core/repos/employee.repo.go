package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"

	"gorm.io/gorm"
)

type EmployeeRepo struct{}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{}
}

func (repo *EmployeeRepo) FindId(employeeId *int) (models.Employee, error) {
	var data models.Employee
	result := db.Database.Limit(1).Find(&data, *employeeId)
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return data, nil
}

func (repo *EmployeeRepo) UpdateById(employee *models.Employee) (int64, error) {
	var effectedRows int64 = 0
	var err error
	return effectedRows, db.Database.Transaction(func(tx *gorm.DB) error {
		err =  tx.Model(employee).Association("Courses").Replace(employee.Courses)
		if err != nil {
			return err
		}

		err =  tx.Model(employee).Association("Degrees").Replace(employee.Degrees)
		if err != nil {
			return err
		}
		result := tx.Updates(employee)

		effectedRows = result.RowsAffected
		err = result.Error
		return result.Error
	})
}

func (repo *EmployeeRepo) Create(newEmployee *models.Employee) error {
	result := db.Database.Create(newEmployee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) Delete(employeeId *int) error {
	result := db.Database.Delete(&models.Employee{}, *employeeId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) GetOneById(employeeId *int) (*models.Employee, error) {
	var data models.Employee
	result := db.Database.Limit(1).Find(&data, *employeeId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *EmployeeRepo) GetOneByName(name string) (*models.Employee, error) {
	var data models.Employee
	result := db.Database.Where("name = ?", name).Limit(1).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *EmployeeRepo) List(pageOpt *dtos.PageOpt, dto *dtos.EmployeeFilter) (*types.ListData[models.Employee], error) {
	query := db.Database.Preload("Degrees").Preload("Courses")
	if dto.EmployeeName != "" {
		query = query.Where(`name LIKE ?`, "%"+dto.EmployeeName+"%")
	}
	if len(dto.EmployeeType) > 0 {
		query = query.Where(`employees.employee_type IN ?`, dto.EmployeeType)
	}
	if dto.StartSalary != 0 || dto.EndSalary != 0 {
		if dto.StartSalary != 0 && dto.EndSalary != 0 {
			query = query.Where("employees.salary BETWEEN ? AND ?", dto.StartSalary, dto.EndSalary)
		} else if dto.StartSalary != 0 {
			query = query.Where("employees.salary > ?", dto.StartSalary)
		} else {
			query = query.Where("employees.salary BETWEEN 0 AND ?", dto.EndSalary)
		}
	}
	return List[models.Employee](pageOpt, query, "employees")
}

func (repo *EmployeeRepo) All(dto *dtos.EmployeeFilter) (*[]models.Employee, error) {
	var data []models.Employee
	query := db.Database.Model(&models.Employee{}).Preload("Schedules").
		Joins("LEFT JOIN schedules ON schedules.employee_id = employees.id AND schedules.scope = ?", dto.Scope)

	if len(dto.EmployeeType) > 0 {
		query = query.Where("employees.employee_type IN ?", dto.EmployeeType)
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

func (repo *EmployeeRepo) FindTelegramId(telegramId *int64) (*models.Employee, error) {
	var data models.Employee
	result := db.Database.Where("telegram_id = ?", telegramId).Limit(1).Find(&data)
	return &data, result.Error
}
