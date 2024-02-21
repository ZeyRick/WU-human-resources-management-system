package dtos

type CreateClockSetting struct {
	Coordinate string `json:"coordinate" gorm:"primaryKey;autoIncrement"`
	ClockRange *int   `json:"clockRange" gorm:"type:int;not null"`
	AllowTime  *int   `json:"allowTime" gorm:"type:int"`
}
