package services

import (
	"backend/pkg/https"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type HelloWorldService struct{}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

func (srv *HelloWorldService) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Do some business logic to get the result")
	 res := https.ErrorBody{
		Msg: "Hello",
		Code:  0,
	}
	render.JSON(w, r, res)
}
