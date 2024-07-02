package models

import (
	"backend/core/types"
	"time"
)

type Clock struct {
	BaseModel
	EmployeeId     *int            `json:"employeeId" gorm:"type:int;not null"`
	ClockType      types.ClockType `json:"clockType" gorm:"type:ENUM;not null"`
	ClockInId      *int            `json:"clockInId" gorm:"int"`
	ClockIn        *Clock          `json:"clockIn" gorm:"foreignKey:ClockInId;references:ID"`
	ClockOutMinute *int            `json:"clockOutMinute" gorm:"int"`
	Employee       Employee        `json:"employee"`
	ScheduleId     *int            `json:"scheduleId"`
	Schedule       Schedule        `json:"schedule"`
	Status         string          `json:"status"`
	EarlyMinutes   *int            `json:"earlyMinutes" gorm:"type:int;default 0"`
	LateMinutes    *int            `json:"lateMinutes" gorm:"type:int;default 0"`
	EditedBy       *uint           `json:"editedBy"`
	Editor         *User           `json:"editor" gorm:"foreignKey:EditedBy;references:ID"`
	DegreeId       *int            `json:"degreeId" gorm:"type:int;not null"`
	Degree         Degree          `json:"degree"`
	CourseId       *int            `json:"courseId" gorm:"type:int;not null"`
	Course         Course          `json:"course"`
	ClockTime      time.Time       `json:"clockTime"`
}

type ClockReports struct {
	TotalWorkMinute  *int   `json:"totalWorkMinute" gorm:"type:string;not null" `
	TotalEarlyMinute *int   `json:"totalEarlyMinute" gorm:"type:string;not null"`
	TotalLateMinute  *int   `json:"totalLateMinute" gorm:"type:string;not null"`
	EmployeeId       *int   `json:"employeeId" gorm:"type:string;not null"`
	Name             string `json:"employeeName" gorm:"type:string;not null"`
	DegreeAlias      string `json:"degreeAlias" gorm:"type:string;not null"`
}
