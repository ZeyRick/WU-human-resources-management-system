package jwttoken

import (
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
		return errors.New("Missing ID Claim")
	}
	if c.ExpireAt <= time.Now().Unix() {
		return errors.New("Missing Expires Claim")
	}
	return nil
}

func GenterateToken(userID uint, expiredTime int) (string, error) {
	expires := time.Now().UTC().Add(time.Hour * time.Duration(expiredTime))
	claims := &Claims{
		ID:       userID,
		ExpireAt: expires.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

func SetCookie(w http.ResponseWriter, token string, cookieName string) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
}

func CheckCookie(w http.ResponseWriter, r *http.Request, cookieName string, expiredTime int) (bool, uint) {
	var neededCookie *http.Cookie
	cookies := r.Cookies()
	if len(cookies) == 0 {
		return false, 0
	}
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			neededCookie = cookie
			break
		}
	}
	if neededCookie == nil {
		return false, 0
	}
	ok, userID := ValidateCookie(w, r, neededCookie.Value, neededCookie)
	if ok == false {
		DeleteCookie(w, cookieName)
		return false, 0
	} else {
		err := UpdateCookieExpiredTime(w, cookieName, expiredTime, userID)
		if err != nil {
			return false, 0
		}
		return true, userID
	}
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
	if claims.ExpireAt <= time.Now().Unix() {
		return false, 0
	}
	return true, claims.ID
}

func DeleteCookie(w http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:   cookieName,
		MaxAge: -1,
		Path:   "/",
		HttpOnly: false,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
}

func UpdateCookieExpiredTime(w http.ResponseWriter, cookieName string, expiredTime int, id uint) error {
	token, err := GenterateToken(id, expiredTime)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
	}
	http.SetCookie(w, cookie)
	return nil
}
