package services

import (
	"backend/adapters/dtos"
	"backend/core/models/user"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/hush"
	"backend/pkg/jwttoken"
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
	if err != nil {
		helper.UnexpectedError(w,r,err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Register complete")
}

func (srv *UserService) UserLogin(w http.ResponseWriter, r *http.Request, payload *dtos.UserLogin) {
	newuser, err := srv.usermodel.FindByUserName(payload.Username)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if newuser.Username == "" || hush.ComparePassword(newuser.Password, payload.Password) != nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Wrong username or password")
		return
	}
	token, err := jwttoken.GenterateToken(newuser.ID, 24 * 7)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, token)	
}

func (srv *UserService) GetUserData(params *dtos.ListUser) (*[]user.User, error) {
	result, err := srv.usermodel.All(params)
	return result, err
}
