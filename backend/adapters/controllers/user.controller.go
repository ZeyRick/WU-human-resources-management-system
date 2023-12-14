package controllers

import (
	"backend/core/services"
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
	ctrl.userservice.UserRegister(w, r)
	return
}

func (ctrl *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	ctrl.userservice.UserLogin(w, r)
	return
}
