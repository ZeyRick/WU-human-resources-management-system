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
	ID                uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeId        int               `json:"employeeId" gorm:"type:int;not null"`
	Scope             string            `json:"scope" gorm:"type:string;not null"`
	Dates             string            `json:"dates" gorm:"tyope:string"`
	ClockInTime       time.Time         `json:"clockInTime"`
	ClockOutTime      time.Time         `json:"clockOutTime"`
	MinuteWorkPerDay  *int              `json:"minuteWorkPerDay"`
	MinuteBreakPerDay *int              `json:"minuteBreakPerDay"`
	Employee          employee.Employee `json:"employee" gorm:"foreignkey:EmployeeId"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
}

type ScheduleRepo struct{}

func NewScheduleRepo() *ScheduleRepo {
	return &ScheduleRepo{}
}

func (repo *ScheduleRepo) Create(newSchedules *Schedule) error {
	result := db.Database.Create(newSchedules)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ScheduleRepo) Update(employeeIds []int, newSchedule *Schedule) error {
	result := db.Database.Table("schedules").Where("employee_id IN ? AND scope = ?", employeeIds, newSchedule.Scope).
		Updates(newSchedule)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ScheduleRepo) BatchCreate(newSchedule *[]Schedule) error {
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
	query := db.Database.Joins(`JOIN employees ON employees.id = schedules.employee_id`).Preload("Employee")
	
	if dto.EmployeeId != nil && *dto.EmployeeId != 0 {
		query = query.Where(`employees.id = ?`, *dto.EmployeeId)
	}
	result := query.Find(&data, "scope = ?", dto.Scope)
	return &data, result.Error
}


func (repo *ScheduleRepo) GetOneById(id uint) (Schedule, error) {
	var data Schedule
	result := db.Database.Where("id = ?", id).First(&data)
	return data, result.Error
}