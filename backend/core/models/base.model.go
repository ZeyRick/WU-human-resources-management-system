package models

import (
	"backend/adapters/dtos"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func Paginate(pageOpt *dtos.PageOpt) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
	pageSize := *pageOpt.Page
	offset := (*pageOpt.Page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}