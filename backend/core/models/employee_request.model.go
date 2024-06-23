package models

import "time"

type EmployeeRequest struct {
	ID               uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       uint              `json:"employeeId" gorm:"type:string"`
	Employee         Employee `json:"employee"`
	TelegramID       int64             `json:"telegramId" gorm:"type:int;not null"`
	TelegramUsername string            `json:"telegramUsername" gorm:"type:string"`
	CreatedAt        time.Time         `json:"createdAt"`
}