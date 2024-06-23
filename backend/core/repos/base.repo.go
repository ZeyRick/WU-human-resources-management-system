package repos

import (
	"backend/adapters/dtos"
	"backend/core/types"
	"math"

	"gorm.io/gorm"
)

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
	selectResult := db.Scopes(paginate(pageOpt)).Order("id DESC").Find(&data)
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

func CustomList[T any](pageOpt *dtos.PageOpt, db *gorm.DB) (*types.ListData[T], error) {
	var data []T
	selectResult := db.Scopes(paginate(pageOpt)).Scan(&data)
	if selectResult.Error != nil {
		return nil, selectResult.Error
	}
	var count int64
	countResult := db.Count(&count)
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
