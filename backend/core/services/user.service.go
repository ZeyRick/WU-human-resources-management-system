package services

import (
	"backend/core/models/user"
	"backend/pkg/hush"
	"backend/pkg/jwttoken"
	"backend/pkg/logger"
	"net/http"
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
	if err != nil {
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
	err = srv.usermodel.Create(&newuser)
	if err != nil {
		logger.Trace(err)
		w.Write([]byte("Error"))
		return
	}
	w.Write([]byte("Register Complete"))
}

func (srv *UserService) UserLogin(w http.ResponseWriter, r *http.Request) {
	jwttoken.CheckCookie(w, r, "LoginCookie", "Cookie Found", "Cookie Not Found")
	username := r.FormValue("username")
	password := r.FormValue("password")
	newuser, err := srv.usermodel.FindByUserName(username)
	if err != nil {
		logger.Trace(err)
		w.Write([]byte("Error"))
		return
	}
	if newuser.Username == "" {
		w.Write([]byte("Username is Incorrect"))
		return
	}
	if hush.ComparePassword(newuser.Password, password) != nil {
		w.Write([]byte("Password is incorrect"))
		return
	}
	token, err := jwttoken.GenterateToken(newuser)
	if err != nil {
		logger.Trace(err)
		w.Write([]byte("Error"))
		return
	}
	jwttoken.SetCookie(w, token, "LoginCookie")
	w.Write([]byte("Logged in"))
}
