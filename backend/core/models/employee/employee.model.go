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
	Name         string                `json:"name" gorm:"type:string;not null"`
	DepartmentId *int                  `json:"departmentId" gorm:"type:number;not null"`
	Department   department.Department `json:"department"`
	ProfilePic   string                `json:"profilePic" gorm:"type:string;not null"`
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

func (repo *EmployeeRepo) All(dto *dtos.EmployeeFilter) (*[]types.EmployeeWithSchedule, error) {
	var data []types.EmployeeWithSchedule
	query := db.Database.Model(&Employee{}).
		Joins("LEFT JOIN schedules ON schedules.employee_id = employees.id AND schedules.scope = ?", dto.Scope).
		Joins("LEFT JOIN departments ON departments.id = employees.department_id").
		Select(`
		employees.id,
		employees.name,
		employees.department_id,
		employees.profile_pic,
		departments.id AS department_id,
		departments.alias AS department_alias,
		departments.created_at AS department_created_at,
		departments.updated_at AS department_updated_at,
		COALESCE(schedules.id, 0) AS schedule_id,
		COALESCE(schedules.employee_id, 0) AS schedule_employee_id,
		COALESCE(schedules.scope, '') AS schedule_scope,
		COALESCE(schedules.dates, '') AS schedule_dates,
		COALESCE(schedules.clock_in_time, '0001-01-01 00:00:00') AS schedule_clock_in_time,
		COALESCE(schedules.clock_out_time, '0001-01-01 00:00:00') AS schedule_clock_out_time,
		COALESCE(schedules.created_at, '0001-01-01 00:00:00') AS schedule_created_at,
		COALESCE(schedules.updated_at, '0001-01-01 00:00:00') AS schedule_updated_at
	`)
	if dto.DepartmentId != nil && *dto.DepartmentId != 0 {
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
