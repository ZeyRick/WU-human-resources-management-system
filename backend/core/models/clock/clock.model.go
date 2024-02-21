package clock

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/models/schedule"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)

type Clock struct {
	models.BaseModel
	EmployeeId     *int              `json:"employeeId" gorm:"type:int;not null"`
	ClockType      types.ClockType   `json:"clockType" gorm:"type:ENUM;not null"`
	ClockInId      *int              `json:"clockInId" gorm:"int"`
	ClockIn        *Clock            `json:"clockIn" gorm:"foreignKey:ClockInId;references:ID"`
	ClockOutMinute *int              `json:"clockOutMinute" gorm:"int"`
	Employee       employee.Employee `json:"employee"`
	ScheduleId     *int              `json:"scheduleId"`
	Schedule       schedule.Schedule `json:"schedule"`
}

type ClockRepo struct{}

func NewClockRepo() *ClockRepo {
	return &ClockRepo{}
}

func (repo *ClockRepo) Create(newClock *Clock) error {
	result := db.Database.Create(newClock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ClockRepo) LatestClockIn(employeeId *int) (*Clock, error) {
	var data Clock
	result := db.Database.Last(&data, "employee_id = ? AND clock_type = ?", *employeeId, types.ClockIn)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *ClockRepo) List(pageOpt *dtos.PageOpt, dto *dtos.ClockFilter) (*types.ListData[Clock], error) {
	query := db.Database.Joins(`JOIN employees ON employees.id = clocks.employee_id`).Preload("Employee").Order("id DESC")

	if dto.Date != "" {
		startOfDay, err := time.Parse("2006-01-02 15:04:05", dto.Date)
		if err != nil {
			return nil, err
		}
		endOfDay := startOfDay.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

		query = query.Where("clocks.created_at BETWEEN ? AND ?", startOfDay, endOfDay)
	}

	if dto.EmployeeId != 0 {
		query = query.Where("clocks.employee_id = ?", dto.EmployeeId)
	}

	if dto.EmployeeName != "" {
		query = query.Where(`employees.name LIKE ?`, "%"+dto.EmployeeName+"%")
	}

	if dto.DepartmentId != 0 {
		query = query.Where("employees.department_id = ?", dto.DepartmentId)
	}

	// datetime BETWEEN '2024-01-14 00:00:00' AND '2024-01-14 23:59:59'
	return models.List[Clock](pageOpt, query, "clocks")
}

func (repo *ClockRepo) Attendence(pageOpt *dtos.PageOpt, dto *dtos.AttendenceFilter) (*types.ListData[Clock], error) {
	query := db.Database.
		Joins(`JOIN employees ON employees.id = clocks.employee_id`).
		Joins(`JOIN clocks AS clock_in ON clocks.clock_in_id = clock_in.id`).
		Joins(`JOIN schedules ON clocks.schedule_id = schedules.id`).
		Where(`clocks.clock_type = 'out'`).
		Preload("ClockIn").
		Preload("Schedule").
		Preload("Employee").Order("id DESC")

	if dto.EmployeeId != 0 {
		query = query.Where("clocks.employee_id = ?", dto.EmployeeId)
	}

	if dto.EmployeeName != "" {
		query = query.Where(`employees.name LIKE ?`, "%"+dto.EmployeeName+"%")
	}

	if dto.DepartmentId != 0 {
		query = query.Where("employees.department_id = ?", dto.DepartmentId)
	}

	// datetime BETWEEN '2024-01-14 00:00:00' AND '2024-01-14 23:59:59'
	return models.List[Clock](pageOpt, query, "clocks")
}

func (repo *ClockRepo) SumReport(pageOpt *dtos.PageOpt, dto *dtos.ReportFilter) (*[]types.ClockReports, error) {
	query := db.Database.Table("clocks").
		Joins(`JOIN employees ON employees.id = clocks.employee_id`).Preload("Employee").
		Joins(`JOIN departments ON departments.id = employees.department_id`).Preload("Department").
		Order("clocks.employee_id DESC")
	query = query.Group("clocks.employee_id")
	if dto.StartDate != "" {
		startDate, err := time.Parse("2006-01-02 15:04:05", dto.StartDate)
		if err != nil {
			return nil, err
		}
		endDate, err := time.Parse("2006-01-02 15:04:05", dto.EndDate)
		if err != nil {
			return nil, err
		}
		query = query.Where("clocks.created_at >= ? AND clocks.created_at <= ?", startDate, endDate)
	}

	if dto.EmployeeName != "" {
		query = query.Where(`employees.name LIKE ?`, "%"+dto.EmployeeName+"%")
	}

	if dto.DepartmentId != 0 {
		query = query.Where("employees.department_id = ?", dto.DepartmentId)
	}

	query = query.Select(`clocks.employee_id, employees.*, departments.*, SUM(clocks.clock_out_minute) as total_work_minute`)
	var clockReports []types.ClockReports
	result := query.Scan(&clockReports)
	if result.Error != nil {
		return &clockReports, result.Error
	}
	return &clockReports, nil
}

func (repo *ClockRepo) UpdateById(clock *Clock) (int64, error) {
	result := db.Database.Model(&Clock{}).Where("id = ?", clock.ID).Updates(*clock)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
