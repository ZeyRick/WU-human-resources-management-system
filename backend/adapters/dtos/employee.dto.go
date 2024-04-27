package dtos

import "backend/core/types"

type AddEmployee struct {
	Name          string             `json:"name" validate:"required"`
	ProfilePic    string             `json:"profilePic" validate:"omitempty"`
	DepartmentId  int                `json:"departmentId" validate:"required"`
	EmployeeType  types.EmployeeType `json:"employeeType" validate:"required"`
	Salary        float64            `json:"salary" validate:"required"`
	IdNumber      string             `json:"idNumber" validate:"omitempty"`
	IdFileName    string             `json:"idFileName" validate:"omitempty"`
	PhotoFileName string             `json:"photoFileName" validate:"omitempty"`
}

type EmployeeFilter struct {
	EmployeeName string             `json:"employeeName" validate:"omitempty"`
	DepartmentId *int               `json:"departmentId" validate:"omitempty"`
	EmployeeId   *int               `json:"employeeId" validate:"omitempty"`
	EmployeeType types.EmployeeType `json:"employeeType" validate:"omitempty"`
	StartSalary  float64            `json:"startSalary" validate:"omitempty"`
	EndSalary    float64            `json:"endSalary" validate:"omitempty"`
	Scope        string             `json:"scope" validate:"omitempty"`
}
