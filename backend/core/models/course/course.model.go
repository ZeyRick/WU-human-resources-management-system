package course

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
	"time"
)


type Course struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Alias     string    `json:"alias" gorm:"unique;not null"`
    CreatedAt time.Time `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type CourseRepo struct{}

func NewCourseRepo() *CourseRepo {
	return &CourseRepo{}
}

func (repo *CourseRepo) All(dto *dtos.CourseFilter) (*[]Course, error) {
	var data []Course
	query := db.Database
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}

func (repo *CourseRepo) List(pageOpt *dtos.PageOpt, dto *dtos.CourseFilter) (*types.ListData[Course], error) {
	query := db.Database
	if dto.Alias != "" {
		query = query.Where(`alias LIKE ?`, "%"+dto.Alias+"%")
	}
	return models.List[Course](pageOpt, query, "courses")
}

func (repo *CourseRepo) Create(newCourse *Course) error {
	result := db.Database.Create(newCourse)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CourseRepo) UpdateById(employee *Course) (int64, error) {
	result := db.Database.Model(&Course{}).Where("id = ?", employee.ID).Updates(*employee)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}