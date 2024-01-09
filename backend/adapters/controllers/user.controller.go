package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/https"
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
	err = ctrl.userservice.UserRegister(w, &dto)
	if err != nil {
		logger.Trace(err)
		if err.Error() == "409" {
			https.ResponseError(w, r, http.StatusConflict, "Username Already Exist")
		} else {
			https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		}
	}
	https.ResponseMsg(w, r, http.StatusCreated, "Register Complete")
	return
}

func (ctrl *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {

	dto, err := https.GetBody[dtos.UserLogin](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = ctrl.userservice.UserLogin(w, &dto)
	if err != nil {
		logger.Trace(err)
		if err.Error() == "401" {
			https.ResponseError(w, r, http.StatusConflict, "Incorrect Username or Password")
		} else {
			https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		}
		return
	}
	https.ResponseMsg(w, r, http.StatusOK, "Logged in")
	return
}

func (ctrl *UserController) GetUserData(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.GetUserData](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err, users, pageCount := ctrl.userservice.GetUserData(w, r, &dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
	}
	https.ResponseJSON(w, r, http.StatusOK, pageCount)
	for _, user := range users {
		https.ResponseJSON(w, r, http.StatusOK, user)
	}
	return
}
