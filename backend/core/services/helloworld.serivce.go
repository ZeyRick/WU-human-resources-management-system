package services

import (
	"fmt"
	"net/http"
)

type HelloWorldService struct{}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

func (srv *HelloWorldService) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Do some business logic to get the result")
	w.Write([]byte("Return the result for front end"))
}
