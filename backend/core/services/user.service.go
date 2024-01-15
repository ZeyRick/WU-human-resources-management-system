package services

import (
	"backend/adapters/dtos"
	"backend/core/models/user"
	"backend/core/types"
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

func (srv *UserService) UserLogin(w http.ResponseWriter, payload *dtos.UserLogin) (uint, error) {
	newuser, err := srv.usermodel.FindByUserName(payload.Username)
	if err != nil {
		return 0, err
	}
	if newuser.Username == "" || hush.ComparePassword(newuser.Password, payload.Password) != nil {
		err = errors.New("401")
		return 0, err
	}
	token, err := jwttoken.GenterateToken(newuser.ID, 24*30)
	if err != nil {
		return 0, err
	}
	jwttoken.SetCookie(w, token, "LoginCookie")
	return newuser.ID, err
}

func (srv *UserService) GetUserData(params *dtos.ListUser) (*types.ListData[user.User], error) {
	result, err := srv.usermodel.List(params)
	return result, err
}
