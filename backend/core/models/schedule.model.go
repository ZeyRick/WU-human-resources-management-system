package models

import (
	"time"
)

type Schedule struct {
	ID                uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeId        int       `json:"employeeId" gorm:"type:int;not null"`
	Scope             string    `json:"scope" gorm:"type:string;not null"`
	Dates             string    `json:"dates" gorm:"tyope:string"`
	ClockInTime       time.Time `json:"clockInTime"`
	ClockOutTime      time.Time `json:"clockOutTime"`
	MinuteWorkPerDay  *int      `json:"minuteWorkPerDay"`
	MinuteBreakPerDay *int      `json:"minuteBreakPerDay"`
	Employee          Employee  `json:"employee" gorm:"foreignkey:EmployeeId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
