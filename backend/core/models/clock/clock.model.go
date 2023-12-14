package clock

import (
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

func (repo *ClockRepo) FindById(clockId *uint) (Clock, error) {
	clock := Clock{}
	result := db.Database.Limit(1).Find(&clock, *clockId)
	if result.Error != nil {
		return Clock{}, result.Error
	}
	return clock, nil
}

func (repo *ClockRepo) DeleteById(clockId *uint) (int64, error) {
	result := db.Database.Delete(&Clock{}, *clockId)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *ClockRepo) UpdateById(clock *Clock) (int64, error) {
	result := db.Database.Model(&Clock{}).Where("id = ?", clock.ID).Updates(*clock)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *ClockRepo) FindByClockName(clockName string) (Clock, error) {
	clock := Clock{}
	result := db.Database.Where("clockname = ?", clockName).Limit(1).Find(&clock)
	if result.Error != nil {
		return Clock{}, result.Error
	}
	return clock, nil
}
