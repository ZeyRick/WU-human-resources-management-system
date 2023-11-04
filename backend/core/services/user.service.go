package services

import (
	"fmt"
	"net/http"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (srv *HelloWorldService) GetUserName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Do some business logic to get the result")
	w.Write([]byte("Return the result for front end"))
}

func (srv *HelloWorldService) GetUserPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Do some business logic to get the result")
	w.Write([]byte("Return the result for front end"))
}
