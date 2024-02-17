package controllers

import (
	"backend/adapters/dtos"
	"backend/core/services"
	"backend/pkg/helper"
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

func (ctrl *UserController) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	logger.Console("333")
	userId := r.Context().Value("userId").(uint)
	if userId == 0 {
		https.ResponseError(w, r, http.StatusUnauthorized, "Invalid user id")
		return
	}
	user, err := ctrl.userservice.FindById(variable.Create[int](int(userId)))
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if user == nil {
		https.ResponseError(w, r, http.StatusUnauthorized, "Invalid user")
		return
	}
	user.Password = ""
	logger.Console(user)
	https.ResponseJSON(w,r, http.StatusOK, user)
}

func (ctrl *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.UserRegister](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.userservice.UserRegister(w, r, &dto)
}

func (ctrl *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userId, err := https.GetParamsID(r, "userId")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if userId == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing user id")
		return
	}
	user, err := ctrl.userservice.FindById(userId)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if user == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Employee not found")
		return
	}
	ctrl.userservice.Delete(w, r, variable.Create[int](int(user.ID)))
}
func (ctrl *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	dto, err := https.GetBody[dtos.UserLogin](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.userservice.UserLogin(w, r, &dto)
}

func (ctrl *UserController) ResetPW(w http.ResponseWriter, r *http.Request) {
	userId, err := https.GetParamsID(r, "userId")
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if userId == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Missing user id")
		return
	}
	dto, err := https.GetBody[dtos.UserResetPwParams](r)
	if err != nil {
		logger.Trace(err)
		https.ResponseError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	user, err := ctrl.userservice.FindById(userId)
	if err != nil {
		helper.UnexpectedError(w, r, err)
		return
	}
	if user == nil {
		https.ResponseError(w, r, http.StatusBadRequest, "Employee not found")
		return
	}
	ctrl.userservice.ResetPW(w, r, userId, &dto)
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
