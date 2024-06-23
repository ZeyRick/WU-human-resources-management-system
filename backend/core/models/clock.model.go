package models

import (
	"backend/core/types"
	"os/user"
)


type Clock struct {
	BaseModel
	EmployeeId     *int              `json:"employeeId" gorm:"type:int;not null"`
	ClockType      types.ClockType   `json:"clockType" gorm:"type:ENUM;not null"`
	ClockInId      *int              `json:"clockInId" gorm:"int"`
	ClockIn        *Clock            `json:"clockIn" gorm:"foreignKey:ClockInId;references:ID"`
	ClockOutMinute *int              `json:"clockOutMinute" gorm:"int"`
	Employee       Employee `json:"employee"`
	ScheduleId     *int              `json:"scheduleId"`
	Schedule       Schedule `json:"schedule"`
	Status         string            `json:"status"`
	EarlyMinutes   *int              `json:"earlyMinutes" gorm:"type:int;default 0"`
	LateMinutes    *int              `json:"lateMinutes" gorm:"type:int;default 0"`
	EditedBy       *uint             `json:"editedBy"`
	Editor         *user.User        `json:"editor" gorm:"foreignKey:EditedBy;references:ID"`
}
