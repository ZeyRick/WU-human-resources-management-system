package controllers

import (
	"backend/core/services"
	"fmt"
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
	fmt.Println("Parsing data from front end so we can pass to services to process the busines logic")

	ctrl.userservice.UserRegister(w, r)
	return
}
