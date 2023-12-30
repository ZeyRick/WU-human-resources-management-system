package jwttoken

import (
	"backend/core/models/user"
	"backend/pkg/https"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// This is the secret key we need to store it in somewhere safe
var secretkey = "this_is_the_key_needed_to_store_in_env"

type Claims struct {
	Username string
	ExpireAt int64
}

// Valid implements jwt.Claims.
func (c Claims) Valid() error {
	if c.Username == "" {
		return errors.New("missing username claim")
	}
	if c.ExpireAt <= time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func GenterateToken(user user.User) (string, error) {
	claims := &Claims{
		Username: user.Username,
		ExpireAt: time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(secretkey))
}

func SetCookie(w http.ResponseWriter, token string, cookieName string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}

func CheckCookie(w http.ResponseWriter, r *http.Request, cookieName string) bool {
	var neededCookie *http.Cookie
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			neededCookie = cookie
			break
		}
	}
	if cookies != nil {
		if ValidateCookie(w, r, neededCookie.Value) == false {
			https.ResponseText(w, r, 0, "Cookie Not Found")
			return false
		} else {
			https.ResponseText(w, r, 1, "Cookie Found")
			return true
		}
	}
	https.ResponseText(w, r, 0, "Cookie Not Found")
	return false
}

func ValidateCookie(w http.ResponseWriter, r *http.Request, tokenString string) bool {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)
	if err != nil {
		https.ResponseText(w, r, 0, "Invalid Cookie")
		return false
	}
	if claims.ExpireAt <= time.Now().Unix() {
		https.ResponseText(w, r, 0, "Cookie Expired")
		return false
	}
	return true
}
