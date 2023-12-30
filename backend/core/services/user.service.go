package services

import (
	"backend/core/models/user"
	"backend/pkg/db"
	"backend/pkg/https"
	"backend/pkg/hush"
	"backend/pkg/jwttoken"
	"backend/pkg/logger"
	"net/http"
	"strconv"
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
		https.ResponseText(w, r, 0, "Error")
		return
	}
	if newuser.Username != "" {
		https.ResponseText(w, r, 0, "Username is Already Exist")
		return
	}
	password, err = hush.Hush(password)
	if err != nil {
		logger.Trace(err)
		https.ResponseText(w, r, 0, "Error")
		return
	}
	newuser = user.User{Username: username, Name: name, Password: password, ProfilePic: profilePic}
	err = srv.usermodel.Create(&newuser)
	if err != nil {
		logger.Trace(err)
		https.ResponseText(w, r, 0, "Error")
		return
	}
	https.ResponseText(w, r, 1, "Register Complete")
}

func (srv *UserService) UserLogin(w http.ResponseWriter, r *http.Request) {
	//Check if user already has login cookie
	if jwttoken.CheckCookie(w, r, "LoginCookie") {
		https.ResponseText(w, r, 1, "Logged In")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	newuser, err := srv.usermodel.FindByUserName(username)
	if err != nil {
		logger.Trace(err)
		https.ResponseText(w, r, 0, "Error")
		return
	}
	if newuser.Username == "" {
		https.ResponseText(w, r, 0, "Username is Incorrect")
		return
	}
	if hush.ComparePassword(newuser.Password, password) != nil {
		https.ResponseText(w, r, 0, "Password is Incorrect")
		return
	}
	token, err := jwttoken.GenterateToken(newuser)
	if err != nil {
		logger.Trace(err)
		https.ResponseText(w, r, 0, "Error")
		return
	}
	jwttoken.SetCookie(w, token, "LoginCookie")
	https.ResponseText(w, r, 1, "Logged In")
}

func (srv *UserService) GetUserData(w http.ResponseWriter, r *http.Request) {
	dataPerPage, err := strconv.Atoi(r.FormValue("dataPerPage"))
	if err != nil {
		https.ResponseText(w, r, 0, "Incorrect Value Format")
	}
	pageNumber, err := strconv.Atoi(r.FormValue("pageNumber"))
	if err != nil {
		https.ResponseText(w, r, 0, "Incorrect Value Format")
	}
	var count int64
	db.Database.Table("hr_management").Count(&count)
	if err := db.Database.Error; err != nil {
		https.ResponseText(w, r, 0, "Error")
	}
	var pageCount int = int(count) / pageNumber
	offSet := (pageNumber - 1) * dataPerPage
	users, err := srv.usermodel.GetUsers(offSet, dataPerPage)
	pageCount = pageCount
	users = users
}
