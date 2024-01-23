package user

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/pkg/db"
)

type User struct {
	models.BaseModel
	Username   string `json:"username" gorm:"type:string;not null"`
	Name       string `json:"name" gorm:"type:string;not null"`
	Password   string `json:"password" gorm:"type:string;not null"`
	ProfilePic string `json:"profilePic" gorm:"type:string;not null"`
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) Create(newUser *User) error {
	result := db.Database.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepo) FindById(userId *uint) (User, error) {
	user := User{}
	result := db.Database.Limit(1).Find(&user, *userId)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepo) DeleteById(userId *uint) (int64, error) {
	result := db.Database.Delete(&User{}, *userId)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *UserRepo) UpdateById(user *User) (int64, error) {
	result := db.Database.Model(&User{}).Where("id = ?", user.ID).Updates(*user)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (repo *UserRepo) FindByUserName(userName string) (User, error) {
	user := User{}
	result := db.Database.Where("username = ?", userName).Limit(1).Find(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepo) GetUsers(offSet, limit int) ([]User, error) {
	user := []User{}
	err := db.Database.Offset(offSet).Limit(limit).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) All(dto *dtos.ListUser) (*[]User, error) {
	var data []User
	query := db.Database.Where("username != ?", "root").Order("id DESC")
	if (dto.Name != "") {
		query.Where("name = ? AND", dto.Name)
	}
	result := query.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}
