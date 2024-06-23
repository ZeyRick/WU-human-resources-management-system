package models

import "time"


type ClockSetting struct {
	Coordinate string    `json:"coordinate" gorm:"primaryKey;autoIncrement"`
	ClockRange *int      `json:"clockRange" gorm:"type:int;default:50"`
	AllowTime  *int      `json:"allowTime" gorm:"type:int;default:0"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
