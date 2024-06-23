package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
)

type DegreeRepo struct{}

func NewDegreeRepo() *DegreeRepo {
	return &DegreeRepo{}
}
func (repo *DegreeRepo) All() (*[]models.Degree, error) {
	var data []models.Degree
	query := db.Database
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}

func (repo *DegreeRepo) List(pageOpt *dtos.PageOpt, dto *dtos.DegreeFilter) (*types.ListData[models.Degree], error) {
	query := db.Database
	if dto.Alias != "" {
		query = query.Where(`alias LIKE ?`, "%"+dto.Alias+"%")
	}
	return List[models.Degree](pageOpt, query, "degrees")
}

func (repo *DegreeRepo) Create(newDegree *models.Degree) error {
	result := db.Database.Create(newDegree)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *DegreeRepo) UpdateById(degree *models.Degree) (int64, error) {
	result := db.Database.Model(&models.Degree{}).Where("id = ?", degree.ID).Updates(*degree)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
