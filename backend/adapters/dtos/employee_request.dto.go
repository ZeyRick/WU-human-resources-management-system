package dtos

import "backend/core/types"

type Confirmation struct {
	TelegramID   *int64                 `json:"telegramID" validate:"required"`
	Confirmation types.ConfirmationType `json:"confirmation" validate:"required"`
}

type EmployeeRequestFilter struct {
	EmployeeID       int    `json:"employeeId" validate:"omitempty"`
	TelegramUsername string `json:"telegramUsername" validate:"omitempty"`
}
