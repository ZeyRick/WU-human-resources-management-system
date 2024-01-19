package models

import (
	"backend/adapters/dtos"
	"backend/core/types"
	"math"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
}

func paginate(pageOpt *dtos.PageOpt) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
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

func List[T any](pageOpt *dtos.PageOpt, db *gorm.DB, tableName string) (*types.ListData[T], error) {
	var data []T
	selectResult := db.Scopes(paginate(pageOpt)).Find(&data)
	if selectResult.Error != nil {
		return nil, selectResult.Error
	}
	var count int64
	countResult := db.Table(tableName).Count(&count)
	if countResult.Error != nil {
		return nil, countResult.Error
	}
	totalPage := int64(math.Ceil(float64(count) / float64(*pageOpt.Size)))
	pageOptFinal := types.Pagination{
		PageSize:   pageOpt.Size,
		CurPage:    pageOpt.Page,
		TotalPage:  &totalPage,
		TotalCount: &count,
	}
	return &types.ListData[T]{PageOpt: &pageOptFinal, Data: &data}, nil
}
