package dtos

import "backend/core/types"

type Confirmation struct {
	RequestID    *int                   `json:"requestId" validate:"required"`
	Confirmation types.ConfirmationType `json:"confirmation" validate:"required"`
}

type EmployeeRequestFilter struct {
	EmployeeID       int    `json:"employeeId" validate:"omitempty"`
	CourseId     int    `json:"courseId" validate:"omitempty"`
	EmployeeName     string `json:"employeeName" validate:"omitempty"`
	TelegramUsername string `json:"telegramUsername" validate:"omitempty"`
	TelegramId       string `json:"telegramId" validate:"omitempty"`
}
