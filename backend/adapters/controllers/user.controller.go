package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
	"backend/pkg/https"
	"backend/pkg/jwttoken"
	"backend/pkg/logger"
	"net/http"
)

type UserController struct {
	userservice *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userservice: services.NewUserService(),
	}
}

func (ctrl *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.UserRegister](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.userservice.UserRegister(w,r, &dto)
}

func (ctrl *UserController) UserLogout(w http.ResponseWriter, r *http.Request) {
	jwttoken.DeleteCookie(w, "LoginCookie")
}

func (ctrl *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.UserLogin](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = ctrl.userservice.UserLogin(w, r, &dto)
	if err != nil {
		if err.Error() == "401" {
			https.ResponseError(w, r, http.StatusBadRequest, "Incorrect Username or Password")
			return
		} else {
			logger.Trace(err)
			https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
			return
		}
	}
}

func (ctrl *UserController) GetUserData(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.ListUser](r)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	result, err := ctrl.userservice.GetUserData(&dto)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
}
