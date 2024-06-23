package models

import "time"

type Course struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Alias     string    `json:"alias" gorm:"unique;not null"`
    CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
