package controllers

import (
	"backend/core/services"
	"fmt"
	"net/http"
)

type UserController struct {
	helloWorldService *services.HelloWorldService
}

func NewUserController() *HelloWorldController {
	return &HelloWorldController{
		helloWorldService: services.NewHelloWorldService(),
	}
}

func (ctrl *UserController) GetUserName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parsing data from front end so we can pass to services to process the busines logic")

	ctrl.helloWorldService.GetUserName(w, r)
	return
}

func (ctrl *UserController) GetUserPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parsing data from front end so we can pass to services to process the busines logic")

	ctrl.helloWorldService.GetUserPassword(w, r)
	return
}
