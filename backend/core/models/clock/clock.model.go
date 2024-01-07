package clock

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/models/employee"
	"backend/core/types"
	"backend/pkg/db"
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

func (repo *ClockRepo) List(dto *dtos.ListClock) (*types.ListData[Clock], error) {
	return models.List[Clock](dto.PageOpt, db.Database, "clocks")
}

func (repo *ClockRepo) UpdateById(clock *Clock) (int64, error) {
	result := db.Database.Model(&Clock{}).Where("id = ?", clock.ID).Updates(*clock)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
