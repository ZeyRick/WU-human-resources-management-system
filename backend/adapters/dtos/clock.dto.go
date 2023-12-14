package dtos

import "backend/core/types"

type Clock struct {
	EmployeeId *int            `json:"employeeId"`
	ClockType  types.ClockType `json:"clockType"`
}
