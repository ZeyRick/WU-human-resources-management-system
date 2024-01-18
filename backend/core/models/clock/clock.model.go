package clock

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)

type Clock struct {
	models.BaseModel
	EmployeeId *int            `gorm:"type:int;not null"`
	ClockType  types.ClockType `gorm:"type:ENUM;not null"`
	Employee   employee.Employee
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

func (repo *ClockRepo) List(pageOpt *dtos.PageOpt, dto *dtos.ClockFilter) (*types.ListData[Clock], error) {
	query := db.Database.Joins("Employee")

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

	// datetime BETWEEN '2024-01-14 00:00:00' AND '2024-01-14 23:59:59'
	return models.List[Clock](pageOpt, query, "clocks")
}

func (repo *ClockRepo) UpdateById(clock *Clock) (int64, error) {
	result := db.Database.Model(&Clock{}).Where("id = ?", clock.ID).Updates(*clock)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
