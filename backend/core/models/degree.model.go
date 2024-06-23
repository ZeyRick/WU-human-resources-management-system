package models

import "time"

type Degree struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Alias     string    `json:"alias" gorm:"unique;not null"`
	Rate      float64   `json:"rate" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}