package schedule

import (
	"backend/core/models/employee"
	"backend/pkg/db"
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