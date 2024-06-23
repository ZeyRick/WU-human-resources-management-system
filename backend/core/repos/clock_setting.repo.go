package repos

import (
	"backend/core/models"
	"backend/pkg/db"

	"gorm.io/gorm"
)

type ClockSettingRepo struct{}

func NewClockSettingRepo() *ClockSettingRepo {
	return &ClockSettingRepo{}
}

func (repo *ClockSettingRepo) Update(newClockSetting *models.ClockSetting) (*int64, error) {
	result := db.Database.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(models.ClockSetting{}).Limit(1).Updates(newClockSetting)
	return &result.RowsAffected, result.Error
}

func (repo *ClockSettingRepo) Create(newClockSetting *models.ClockSetting) error {
	result := db.Database.Create(newClockSetting)
	return result.Error
}

func (repo *ClockSettingRepo) Get() (*models.ClockSetting, error) {
	var data models.ClockSetting
	result := db.Database.Limit(1).Find(&data)
	return &data, result.Error
}
