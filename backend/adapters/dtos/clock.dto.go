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
	EmployeeName string `json:"employeeName,string,omitempty"`
	EmployeeId   int    `json:"employeeId,string,omitempty"`
	DepartmentId int    `json:"departmentId,string,omitempty"`
}
