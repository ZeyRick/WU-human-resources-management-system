package department

import (
	"backend/adapters/dtos"
	"backend/pkg/db"
	"time"
)

type Department struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Alias     string `json:"alias" gorm:"type:string;not null"`
	CreatedAt time.Time `json:"createdAt"` 
	UpdatedAt time.Time	`json:"updatedAt"`
}

type DepartmentRepo struct{}

func NewDepartmentRepo() *DepartmentRepo {
	return &DepartmentRepo{}
}

func (repo *DepartmentRepo) All(dto *dtos.DepartmentFilter) (*[]Department, error) {
	var data []Department
	query := db.Database
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}

