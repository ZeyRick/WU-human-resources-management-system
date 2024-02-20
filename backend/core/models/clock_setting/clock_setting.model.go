package clocksetting

import (
	"backend/pkg/db"
	"time"

	"gorm.io/gorm"
)

type ClockSetting struct {
	Coordinate string    `json:"coordinate" gorm:"primaryKey;autoIncrement"`
	ClockRange *int      `json:"clockRange" gorm:"type:int;default:50"`
	AllowTime  *int      `json:"allowTime" gorm:"type:int;default:0"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ClockSettingRepo struct{}

func NewClockSettingRepo() *ClockSettingRepo {
	return &ClockSettingRepo{}
}

func (repo *ClockSettingRepo) Update(newClockSetting *ClockSetting) (*int64, error) {
	result := db.Database.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(ClockSetting{}).Limit(1).Updates(newClockSetting)
	return &result.RowsAffected, result.Error
}

func (repo *ClockSettingRepo) Create(newClockSetting *ClockSetting) error {
	result := db.Database.Create(newClockSetting)
	return result.Error
}

func (repo *ClockSettingRepo) Get() (*ClockSetting, error) {
	var data ClockSetting
	result := db.Database.Limit(1).Find(&data)
	return &data, result.Error
}
