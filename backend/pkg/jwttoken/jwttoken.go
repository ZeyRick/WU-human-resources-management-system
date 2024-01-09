package jwttoken

import (
	"backend/core/models/user"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// This is the secret key we need to store it in somewhere safe
var secretkey = "this_is_the_key_needed_to_store_in_env"

type Claims struct {
	ID       uint
	ExpireAt int64
}

// Valid implements jwt.Claims.
func (c Claims) Valid() error {
	if c.ID == 0 {
		return errors.New("missing username claim")
	}
	if c.ExpireAt <= time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func GenterateToken(user user.User) (string, error) {
	claims := &Claims{
		ID: user.ID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(secretkey))
}

func SetCookie(w http.ResponseWriter, token string, cookieName string, expiredTime int) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * time.Duration(expiredTime)),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}

func CheckCookie(w http.ResponseWriter, r *http.Request, cookieName string, expiredTime int) (bool, uint) {
	var neededCookie *http.Cookie
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			neededCookie = cookie
			break
		}
	}
	if cookies != nil {
		ok, userID := ValidateCookie(w, r, neededCookie.Value, neededCookie)
		if ok == false {
			DeleteCookie(w, cookieName)
			return false, 0
		} else {
			UpdateCookieExpiredTime(w, cookieName, expiredTime)
			return true, userID
		}
	}
	return false, 0
}

func ValidateCookie(w http.ResponseWriter, r *http.Request, tokenString string, cookie *http.Cookie) (bool, uint) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)
	if err != nil {
		return false, 0
	}
	if cookie.Expires.Unix() <= time.Now().Unix() {
		return false, 0
	}
	return true, claims.ID
}

func DeleteCookie(w http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:   cookieName,
		MaxAge: -1,
		Path:   "/",
		Secure: true,
	}
	http.SetCookie(w, cookie)
}

func UpdateCookieExpiredTime(w http.ResponseWriter, cookieName string, expiredTime int) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Expires:  time.Now().Add(time.Hour * time.Duration(expiredTime)),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}
