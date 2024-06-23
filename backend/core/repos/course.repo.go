package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/types"
	"backend/pkg/db"
)


type CourseRepo struct{}

func NewCourseRepo() *CourseRepo {
	return &CourseRepo{}
}

func (repo *CourseRepo) All(dto *dtos.CourseFilter) (*[]models.Course, error) {
	var data []models.Course
	query := db.Database
	dbRes := query.Find(&data)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return &data, nil
}

func (repo *CourseRepo) List(pageOpt *dtos.PageOpt, dto *dtos.CourseFilter) (*types.ListData[models.Course], error) {
	query := db.Database
	if dto.Alias != "" {
		query = query.Where(`alias LIKE ?`, "%"+dto.Alias+"%")
	}
	return List[models.Course](pageOpt, query, "courses")
}

func (repo *CourseRepo) Create(newCourse *models.Course) error {
	result := db.Database.Create(newCourse)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CourseRepo) UpdateById(employee *models.Course) (int64, error) {
	result := db.Database.Model(&models.Course{}).Where("id = ?", employee.ID).Updates(*employee)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *CourseRepo) FindByIds(ids []int) ([]models.Course, error) {
	var result []models.Course
	err := db.Database.Find(&result, ids).Error
	return result, err
}