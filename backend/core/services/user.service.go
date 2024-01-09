package services

import (
	"backend/adapters/dtos"
	"backend/core/models/user"
	"backend/pkg/db"
	"backend/pkg/hush"
	"backend/pkg/jwttoken"
	"errors"
	"net/http"
)

type UserService struct {
	usermodel *user.UserRepo
}

func NewUserModel() *UserService {
	return &UserService{
		usermodel: user.NewUserRepo(),
	}
}

func NewUserService() *UserService {
	return &UserService{}
}

func (srv *UserService) UserRegister(w http.ResponseWriter, payload *dtos.UserRegister) error {
	newuser, err := srv.usermodel.FindByUserName(payload.Username)
	if err != nil {
		return err
	}
	if newuser.Username != "" {
		err = errors.New("409")
		return err
	}
	password, err := hush.Hush(payload.Password)
	if err != nil {
		return err
	}
	newuser = user.User{Username: payload.Username, Name: payload.Name, Password: password, ProfilePic: payload.ProfilePic}
	err = srv.usermodel.Create(&newuser)
	return err
}

func (srv *UserService) UserLogin(w http.ResponseWriter, payload *dtos.UserLogin) error {
	newuser, err := srv.usermodel.FindByUserName(payload.Username)
	if err != nil {
		return err
	}
	if newuser.Username == "" || hush.ComparePassword(newuser.Password, payload.Password) != nil {
		err = errors.New("401")
		return err
	}
	token, err := jwttoken.GenterateToken(newuser)
	if err != nil {
		return err
	}
	jwttoken.SetCookie(w, token, "LoginCookie", 24*30)
	return err
}

func (srv *UserService) GetUserData(w http.ResponseWriter, r *http.Request, payload *dtos.GetUserData) (error, []user.User, int) {
	var count int64
	db.Database.Table("hr_management").Count(&count)
	if err := db.Database.Error; err != nil {
		return err, nil, 0
	}
	pageCount := int(count) / int(*payload.PageNumber)
	offSet := (*payload.PageNumber - 1) * *payload.DataPerPage
	users, err := srv.usermodel.GetUsers(int(offSet), int(*payload.DataPerPage))
	return err, users, pageCount
}
