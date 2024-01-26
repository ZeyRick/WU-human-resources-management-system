package services

import (
	"backend/adapters/dtos"
	"backend/core/models/user"
	"backend/pkg/helper"
	"backend/pkg/https"
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

func (srv *UserService) UserRegister(w http.ResponseWriter, r *http.Request, payload *dtos.UserRegister) {
	newuser, err := srv.usermodel.FindByUserName(payload.Username)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if newuser.Username != "" {
		https.ResponseError(w, r, http.StatusBadRequest, "Username Already Exist")
		return
	}
	password, err := hush.Hush(payload.Password)
	if err != nil {
		return
	}
	newuser = user.User{Username: payload.Username, Name: payload.Name, Password: password, ProfilePic: payload.ProfilePic}
	err = srv.usermodel.Create(&newuser)
	https.ResponseMsg(w, r, http.StatusCreated, "Register complete")
}

func (srv *UserService) UserLogin(w http.ResponseWriter, r *http.Request, payload *dtos.UserLogin) error {
	newuser, err := srv.usermodel.FindByUserName(payload.Username)
	if err != nil {
		return err
	}
	if newuser.Username == "" || hush.ComparePassword(newuser.Password, payload.Password) != nil {
		err = errors.New("401")
		return err
	}
	token, err := jwttoken.GenterateToken(newuser.ID, 24*30)
	if err != nil {
		return err
	}

	jwttoken.SetCookie(w, token, "LoginCookie")
	return nil
}

func (srv *UserService) GetUserData(params *dtos.ListUser) (*[]user.User, error) {
	result, err := srv.usermodel.All(params)
	return result, err
}
