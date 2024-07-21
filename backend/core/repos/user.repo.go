package repos

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/pkg/db"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) Create(newUser *models.User) error {
	result := db.Database.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepo) FindById(userId *int) (*models.User, error) {
	user := models.User{}
	result := db.Database.Limit(1).Find(&user, *userId)
	if result.Error != nil {
		return &models.User{}, result.Error
	}
	return &user, nil
}

func (repo *UserRepo) DeleteById(userId *int) (int64, error) {
	result := db.Database.Delete(&models.User{}, *userId)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *UserRepo) UpdateById(user *models.User) (int64, error) {
	result := db.Database.Model(&models.User{}).Where("id = ?", user.ID).Updates(*user)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *UserRepo) FindByUserName(userName string) (models.User, error) {
	user := models.User{}
	result := db.Database.Where("username = ?", userName).Limit(1).Find(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepo) GetUsers(offSet, limit int) ([]models.User, error) {
	user := []models.User{}
	err := db.Database.Offset(offSet).Limit(limit).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) All(dto *dtos.ListUser) (*[]models.User, error) {
	var data []models.User
	query := db.Database.Where("username != ?", "root").Order("id DESC")
	if dto.Name != "" {
		query.Where("name = ? AND", dto.Name)
	}
	result := query.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (repo *UserRepo) Count() (int64, error) {
	var result int64
	err := db.Database.Table("users").Count(&result).Error
	return result, err
}