package models

import "backend/core/types"

type Employee struct {
	BaseModel
	Name             string             `json:"name" gorm:"type:string;not null"`
	TelegramID       int64              `json:"telegramId" gorm:"type:int;not null"`
	TelegramUsername string             `json:"telegramUsername"`
	EmployeeType     types.EmployeeType `json:"employeeType"`
	Salary           float64            `json:"salary"`
	IdNumber         string             `json:"idNumber"`
	IdFileName       string             `json:"idFileName"`
	PhotoFileName    string             `json:"photoFileName"`
	Degrees          []Degree           `json:"courses" gorm:"many2many:employee_degrees;"`
	Courses          []Course           `json:"degrees" gorm:"many2many:employee_courses;"`
}
