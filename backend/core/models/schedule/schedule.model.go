package schedule

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/db"
	"backend/pkg/logger"
	"time"
)

type Schedule struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	EmployeeId *int   `gorm:"type:int;not null"`
	Scope      string `gorm:"type:string;not null"`
	Dates      string `gorm:"tyope:string"`
	Employee   employee.Employee
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
	if result.Error != nil {
		return Schedule{}, result.Error
	}
	return data, nil
}

func (repo *ScheduleRepo) List(pageOpt *dtos.PageOpt, dto *dtos.ScheduleFilter) (*types.ListData[Schedule], error) {
	query := db.Database

	logger.Console(dto.Scope)
	if dto.Scope != "" {
		query = query.Where("scope = ?", dto.Scope)
	}
	return models.List[Schedule](pageOpt, query, "schedules")
}
