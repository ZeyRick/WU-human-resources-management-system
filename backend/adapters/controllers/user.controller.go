package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/https"
	"backend/pkg/logger"
	"backend/pkg/variable"
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
	userID, err := ctrl.userservice.UserLogin(w, &dto)
	if err != nil {
		logger.Trace(err)
		if err.Error() == "401" {
			https.ResponseError(w, r, http.StatusConflict, "Incorrect Username or Password")
		} else {
			https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
		}
		return
	}
	https.ResponseJSON(w, r, http.StatusOK, userID)
	return
}

func (ctrl *UserController) GetUserData(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetQuery[dtos.ListUser](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if dto.PageOpt.Page == nil || *dto.PageOpt.Page == 0 {
		dto.PageOpt.Page = variable.Create[int64](1)
	}
	if dto.PageOpt.Size == nil || *dto.PageOpt.Size == 0 {
		dto.PageOpt.Size = variable.Create[int64](10)
	}
	result, err := ctrl.userservice.GetUserData(&dto)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusInternalServerError, "Something went wrong")
	}
	https.ResponseJSON(w, r, http.StatusOK, *result)
	return
}
