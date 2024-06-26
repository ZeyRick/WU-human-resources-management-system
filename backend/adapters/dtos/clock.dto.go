package dtos

import (
	"backend/core/types"
	"time"
)

type Clock struct {
	EmployeeId *int            `json:"employeeId,string" validate:"required"`
	ClockType  types.ClockType `json:"clockType" validate:"required"`
}

type ListClock struct {
	PageOpt PageOpt `json:"pageOpt" validate:"required"`
}

type ClockFilter struct {
	Date         string `json:"string,omitempty"`
	StartDate    string `json:"startDate,string,omitempty"`
	EndDate      string `json:"endDate,string,omitempty"`
	EmployeeName string `json:"employeeName,string,omitempty"`
	EmployeeId   int    `json:"employeeId,string,omitempty"`
	CourseId     int    `json:"courseId,string,omitempty"`
}

type UpdateClock struct {
	ClockTime string `json:"clockTime" validate:"required"`
}

type AttendenceFilter struct {
	StartDate    string `json:"startDate,string,omitempty"`
	EndDate      string `json:"endDate,string,omitempty"`
	EmployeeName string `json:"employeeName,string,omitempty"`
	EmployeeId   int    `json:"employeeId,string,omitempty"`
	CourseId     int    `json:"courseId,string,omitempty"`
}

type ManualClock struct {
	EmployeeId *int            `json:"employeeId,string" validate:"required"`
	ClockType  types.ClockType `json:"clockType" validate:"required"`
	Degree     string          `json:"degree,string" validate:"required"`
	Course     string          `json:"course,string" validate:"required"`
	ClockTime  time.Time       `json:"clockTime" validate:"required"`
}
