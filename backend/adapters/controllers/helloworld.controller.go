package controllers

import (
	"backend/core/services"
	"fmt"
	"net/http"
)

type HelloWorldController struct {
	helloWorldService *services.HelloWorldService
}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{
		helloWorldService: services.NewHelloWorldService(),
	}
}

func (ctrl *HelloWorldController) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parsing data from front end so we can pass to services to process the busines logic")
	ctrl.helloWorldService.GetHelloWorld(w, r)
	return
}
