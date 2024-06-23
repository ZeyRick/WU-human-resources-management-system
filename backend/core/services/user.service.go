package services

import (
	"backend/adapters/dtos"
	"backend/core/models"
	"backend/core/repos"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/hush"
	"backend/pkg/jwttoken"
	"net/http"
)

type UserService struct {
	usermodel *repos.UserRepo
}

func NewUserModel() *UserService {
	return &UserService{
		usermodel: repos.NewUserRepo(),
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
	newuser = models.User{Username: payload.Username, Name: payload.Name, Password: password, UserLevel: payload.UserLeval}
	err = srv.usermodel.Create(&newuser)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Register complete")
}

func (srv *UserService) Delete(w http.ResponseWriter, r *http.Request, userId *int) {
	_, err := srv.usermodel.DeleteById(userId)
	if err != nil {
		helper.UnexpectedError(w, r, err)
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
	token, err := jwttoken.GenterateToken(newuser.ID, 24*7)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, dtos.UserToken{
		Token: token,
	})
}

func (srv *UserService) ResetPW(w http.ResponseWriter, r *http.Request, userId *int, payload *dtos.UserResetPwParams) {
	curUser, err := srv.usermodel.FindById(userId)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if curUser == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "User not found")
		return
	}
	newPw, err := hush.Hush(payload.Password)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	curUser.Password = newPw
	_, err = srv.usermodel.UpdateById(curUser)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseMsg(w, r, http.StatusOK, "Password reseted")
}

func (srv *UserService) GetUserData(params *dtos.ListUser) (*[]models.User, error) {
	result, err := srv.usermodel.All(params)
	return result, err
}

func (srv *UserService) FindById(userId *int) (*models.User, error) {
	return srv.usermodel.FindById(userId)
}
