package degree

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)

type Degree struct {
	ID        uint      `gorm:"primaryKey"`
	Alias     string    `gorm:"unique;not null"`
	Rate      *float64  `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
type DegreeRepo struct{}

func NewDegreeRepo() *DegreeRepo {
	return &DegreeRepo{}
}
func (repo *DegreeRepo) All() (*[]Degree, error) {
	var data []Degree
	query := db.Database
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}

func (repo *DegreeRepo) List(pageOpt *dtos.PageOpt, dto *dtos.DegreeFilter) (*types.ListData[Degree], error) {
	query := db.Database
	if dto.Alias != "" {
		query = query.Where(`alias LIKE ?`, "%"+dto.Alias+"%")
	}
	return models.List[Degree](pageOpt, query, "degrees")
}

func (repo *DegreeRepo) Create(newDegree *Degree) error {
	result := db.Database.Create(newDegree)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *DegreeRepo) UpdateById(degree *Degree) (int64, error) {
	result := db.Database.Model(&Degree{}).Where("id = ?", degree.ID).Updates(*degree)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
