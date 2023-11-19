package user

import (
	"backend/core/models"
	"backend/pkg/db"
)

type User struct {
	models.BaseModel
	Username   string `gorm:"type:string;not null"`
	Name       string `gorm:"type:string;not null"`
	Password   string `gorm:"type:string;not null"`
	ProfilePic string `gorm:"type:string;not null"`
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
	result := db.Database.Where("username = ?", userName).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
