package dtos

import (
	"backend/core/types"
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
	StartDate    string   `json:"startDate,string,omitempty"`
	EndDate      string   `json:"endDate,string,omitempty"`
	EmployeeName string   `json:"employeeName,string,omitempty"`
	EmployeeId   int      `json:"employeeId,string,omitempty"`
	EmployeeType []string `json:"employeeType,omitempty"`
	IsTeaching   bool     `json:"isTeaching,omitempty"`
	CourseId     int      `json:"courseId,string,omitempty"`
}

type ManualClock struct {
	EmployeeId   *int   `json:"employeeId" validate:"required"`
	DegreeId     int    `json:"degreeId" validate:"required"`
	CourseId     int    `json:"courseId" validate:"required"`
	ClockInTime  string `json:"clockInTime" validate:"required"`
	ClockOutTime string `json:"clockOutTime" validate:"required"`
	TotalMinute  int    `json:"totalMinute" validate:"required"`
}

type UpdateManualClock struct {
	ClockInTime  string `json:"clockInTime" validate:"required"`
	ClockOutTime string `json:"clockOutTime" validate:"required"`
	DegreeId     int    `json:"degreeId" validate:"required"`
	CourseId     int    `json:"courseId" validate:"required"`
}
