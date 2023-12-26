package employee

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
	"math"
)

type Employee struct {
	models.BaseModel
	Name       string `gorm:"type:string;not null"`
	ProfilePic string `gorm:"type:string;not null"`
}

type EmployeeRepo struct{}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{}
}

func (repo *EmployeeRepo) Create(newEmployee *Employee) error {
	result := db.Database.Create(newEmployee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *EmployeeRepo) List(dto *dtos.ListEmployee)(*types.ListData[Employee], error) {
	var data []Employee
	selectResult := db.Database.Scopes(models.Paginate(&dto.PageOpt)).Find(&data)
	if selectResult.Error != nil {
		return nil, selectResult.Error
	}
	var count int64
	countResult := db.Database.Table("employees").Count(&count)
		if countResult.Error != nil {
		return nil, countResult.Error
	}
	totalPage := int64(math.Ceil(float64(count) / float64(*dto.PageOpt.Size)))
	pageOpt := types.Pagination{
		PageSize: dto.PageOpt.Size,
		CurPage: dto.PageOpt.Page,
		TotalPage: &totalPage,
		TotalCount: &count,
	}
	return &types.ListData[Employee]{PageOpt: &pageOpt, Data: &data}, nil
}