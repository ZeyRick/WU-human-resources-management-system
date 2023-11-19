package services

import (
	"backend/core/models/user"
	"backend/pkg/hush"
	"backend/pkg/logger"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type UserService struct {
	usermodel *user.UserRepo
}

func NewUserModel() *UserService {
	return &UserService{
		usermodel: user.NewUserRepo(),
	}
}

func NewUserService() *UserService {
	return &UserService{}
}

func (srv *UserService) UserRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	name := r.FormValue("name")
	password := r.FormValue("password")
	profilePic := r.FormValue("profilepic")
	newuser, err := srv.usermodel.FindByUserName(username)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Trace(err)
		w.Write([]byte("Error"))
		return
	}
	if newuser.Username != "" {
		w.Write([]byte("User Name Already exist"))
		return
	}
	password, err = hush.Hush(password)
	if err != nil {
		logger.Trace(err)
		w.Write([]byte("Error"))
		return
	}
	newuser = user.User{Username: username, Name: name, Password: password, ProfilePic: profilePic}
	fmt.Println(username, password, name, profilePic)
	err = srv.usermodel.Create(&newuser)
	if err != nil {
		logger.Trace(err)
		w.Write([]byte("Error"))
		return
	}
	w.Write([]byte("Register Complete"))
}
