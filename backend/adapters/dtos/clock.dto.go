package dtos

import (
	"backend/core/types"
)

type Clock struct {
	EmployeeId *int            `json:"employeeId,string,required"`
	ClockType  types.ClockType `json:"clockType,required"`
}

type ListClock struct {
	PageOpt PageOpt `json:"pageOpt,required"`
}

type ClockFilter struct {
	Date         string `json:"string,omitempty"`
	EmployeeName string `json:"employeeName,string,omitempty"`
	EmployeeId   int   `json:"employeeId,string,omitempty"`
}
