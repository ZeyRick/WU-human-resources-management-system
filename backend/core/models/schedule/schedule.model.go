package schedule

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)

type Schedule struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	EmployeeId   *int   `gorm:"type:int;not null"`
	Scope        string `gorm:"type:string;not null"`
	Dates        string `gorm:"tyope:string"`
	ClockInTime  time.Time
	ClockOutTime time.Time
	Employee     employee.Employee `gorm:"foreignkey:EmployeeId"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ScheduleRepo struct{}

func NewScheduleRepo() *ScheduleRepo {
	return &ScheduleRepo{}
}

func (repo *ScheduleRepo) Create(newSchedule *Schedule) error {
	result := db.Database.Create(newSchedule)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ScheduleRepo) FindExistedScope(employeeId *int, scope string) (Schedule, error) {
	var data Schedule
	result := db.Database.Where("employee_id = ? AND scope = ?", *employeeId, scope).Limit(1).Find(&data)
	return data, result.Error
}

func (repo *ScheduleRepo) List(pageOpt *dtos.PageOpt, dto *dtos.ScheduleFilter) (*types.ListData[Schedule], error) {
	query := db.Database
	if dto.Scope != "" {
		query = query.Where("scope = ?", dto.Scope)
	}
	return models.List[Schedule](pageOpt, query, "schedules")
}

func (repo *ScheduleRepo) GetAllByScope(dto *dtos.ScheduleFilter) (*[]Schedule, error) {
	var data []Schedule
	query :=  db.Database.Joins(`JOIN employees ON employees.id = schedules.employee_id`).Preload("Employee").
	Where("employees.department_id = ?", *dto.DepartmentId)
	if *dto.EmployeeId != 0 {
		query = query.Where(`employees.id = ?`, *dto.EmployeeId)
	}
	result := query.Find(&data, "scope = ?", dto.Scope)
	return &data, result.Error
}
