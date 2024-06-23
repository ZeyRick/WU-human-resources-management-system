package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
)

type ScheduleRepo struct{}

func NewScheduleRepo() *ScheduleRepo {
	return &ScheduleRepo{}
}

func (repo *ScheduleRepo) Create(newSchedules *models.Schedule) error {
	result := db.Database.Create(newSchedules)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ScheduleRepo) Update(employeeIds []int, newSchedule *models.Schedule) error {
	result := db.Database.Table("schedules").Where("employee_id IN ? AND scope = ?", employeeIds, newSchedule.Scope).
		Updates(newSchedule)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ScheduleRepo) BatchCreate(newSchedule *[]models.Schedule) error {
	result := db.Database.Create(newSchedule)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ScheduleRepo) FindExistedScope(employeeId *int, scope string) (models.Schedule, error) {
	var data models.Schedule
	result := db.Database.Where("employee_id = ? AND scope = ?", *employeeId, scope).Limit(1).Find(&data)
	return data, result.Error
}

func (repo *ScheduleRepo) List(pageOpt *dtos.PageOpt, dto *dtos.ScheduleFilter) (*types.ListData[models.Schedule], error) {
	query := db.Database
	if dto.Scope != "" {
		query = query.Where("scope = ?", dto.Scope)
	}
	return List[models.Schedule](pageOpt, query, "schedules")
}

func (repo *ScheduleRepo) GetAllByScope(dto *dtos.ScheduleFilter) (*[]models.Schedule, error) {
	var data []models.Schedule
	query := db.Database.Joins(`JOIN employees ON employees.id = schedules.employee_id`).Preload("Employee")
	
	if dto.EmployeeId != nil && *dto.EmployeeId != 0 {
		query = query.Where(`employees.id = ?`, *dto.EmployeeId)
	}
	result := query.Find(&data, "scope = ?", dto.Scope)
	return &data, result.Error
}


func (repo *ScheduleRepo) GetOneById(id uint) (models.Schedule, error) {
	var data models.Schedule
	result := db.Database.Where("id = ?", id).First(&data)
	return data, result.Error
}