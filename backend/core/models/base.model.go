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
	var pageSize int64
	var curPage int64
	if pageOpt.Size == nil {
		pageSize = 10
	} else {
		pageSize = *pageOpt.Size
	}
	if pageOpt.Page == nil {
		curPage = 1
	} else {
		curPage = *pageOpt.Page
	}
	offset := (curPage - 1) * pageSize
    return db.Offset(int(offset)).Limit(int(pageSize))
  }
}