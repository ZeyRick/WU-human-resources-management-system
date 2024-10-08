package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
	"backend/pkg/variable"

	"time"

	"gorm.io/gorm"
)

type ClockRepo struct{}

func NewClockRepo() *ClockRepo {
	return &ClockRepo{}
}

func (repo *ClockRepo) Create(newClock *models.Clock) error {
	result := db.Database.Create(newClock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ClockRepo) LatestClockIn(employeeId *int) (*models.Clock, error) {
	var data models.Clock
	result := db.Database.Last(&data, "employee_id = ? AND clock_type = ?", *employeeId, types.ClockIn)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *ClockRepo) LatestClock(employeeId *int) (*models.Clock, error) {
	var data models.Clock
	result := db.Database.Last(&data, "employee_id = ?", *employeeId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *ClockRepo) List(pageOpt *dtos.PageOpt, dto *dtos.ClockFilter) (*types.ListData[models.Clock], error) {
	query := db.Database.Joins(`JOIN employees ON employees.id = clocks.employee_id`).Preload("Employee").Preload("Schedule").Preload("Editor").Order("id DESC")

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

	query = query.Where("employees.employee_type != 'lecture'")

	// datetime BETWEEN '2024-01-14 00:00:00' AND '2024-01-14 23:59:59'
	return List[models.Clock](pageOpt, query, "clocks")
}

func (repo *ClockRepo) Attendence(pageOpt *dtos.PageOpt, dto *dtos.AttendenceFilter) (*types.ListData[models.Clock], error) {
	query := db.Database.
		Joins(`JOIN employees ON employees.id = clocks.employee_id`).
		Where(`clocks.clock_type = 'out'`).
		Preload("ClockIn").
		Preload("Schedule").
		Preload("Employee").Order("id DESC")

	if dto.EmployeeId != 0 {
		query = query.Where("clocks.employee_id = ?", dto.EmployeeId)
	}

	if len(dto.EmployeeType) > 0 {
		query = query.Where(`employees.employee_type IN ?`, dto.EmployeeType)
	}

	if dto.IsTeaching {
		query = query.Where(`clocks.degree_id IS NOT NULL`)
	} else {
		query = query.Where(`clocks.degree_id IS NULL`)
	}

	if dto.EmployeeName != "" {
		query = query.Where(`employees.name LIKE ?`, "%"+dto.EmployeeName+"%")
	}

	if dto.StartDate != "" && dto.EndDate != "" {
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
	// datetime BETWEEN '2024-01-14 00:00:00' AND '2024-01-14 23:59:59'
	return List[models.Clock](pageOpt, query, "clocks")
}

func (repo *ClockRepo) SumReport(pageOpt *dtos.PageOpt, dto *dtos.ReportFilter) (*types.ListData[models.ClockReports], error) {
	query := db.Database.Table("clocks").
		Joins(`JOIN employees ON employees.id = clocks.employee_id`).Preload("Employee").
		Order("clocks.employee_id DESC")
	query = query.Group("clocks.employee_id")
	if dto.StartDate != "" && dto.EndDate != "" {
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

	if len(dto.EmployeeType) > 0 {
		query = query.Where(`employees.employee_type IN ?`, dto.EmployeeType)
	}
	selectStr := `clocks.employee_id, employees.*,
	SUM(clocks.clock_out_minute) as total_work_minute, 
	SUM(clocks.early_minutes) as total_early_minute, 
	SUM(clocks.late_minutes) as total_late_minute`
	if dto.IsTeaching {
		query = query.Joins(`JOIN degrees ON degrees.id = clocks.degree_id`).Preload("Degree")
		query = query.Where(`clocks.degree_id IS NOT NULL`).Group("clocks.degree_id")
		selectStr += `,degrees.alias as degree_alias`
		
	} else {
		query = query.Where(`clocks.degree_id IS NULL`)
	}

	query = query.Select(selectStr)
	return CustomList[models.ClockReports](pageOpt, query)
}

func (repo *ClockRepo) UpdateById(clock *models.Clock) (int64, error) {
	result := db.Database.Model(&models.Clock{}).Where("id = ?", clock.ID).Updates(*clock)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *ClockRepo) GetOneById(id uint) (models.Clock, error) {
	var clock models.Clock
	err := db.Database.First(&clock, id).Error
	return clock, err
}

func (repo *ClockRepo) GetClockOutByClockIn(clockInID uint) (models.Clock, error) {
	var clock models.Clock
	err := db.Database.Preload("Schedule").Where("clock_in_id =?", clockInID).First(&clock).Error
	return clock, err
}

func (repo *ClockRepo) ManualClock(clockIn *models.Clock, clockOut *models.Clock) error {
	return db.Database.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(clockIn).Error
		if err != nil {
			return err
		}
		clockOut.ClockInId = variable.Create(int(clockIn.ID))
		err = tx.Create(clockOut).Error
		return err
	})
}

func (repo *ClockRepo) ManualUpdate(clockIn *models.Clock, clockOut *models.Clock) error {
	return db.Database.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.Clock{}).Where("id = ?", clockIn.ID).Updates(*clockIn).Error
		if err != nil {
			return err
		}
		err = tx.Model(&models.Clock{}).Where("id = ?", clockOut.ID).Updates(*clockOut).Error
		return err
	})
}
