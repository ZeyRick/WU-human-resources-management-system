package dtos

import "backend/core/types"

type Clock struct {
	EmployeeId *int            `json:"employeeId,string,required"`
	ClockType  types.ClockType `json:"clockType,required"`
}


type ListClock struct {
	PageOpt PageOpt `json:"pageOpt,required"`
}