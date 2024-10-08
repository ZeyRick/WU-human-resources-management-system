package jwttoken

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
	secretkey := os.Getenv("JWT_KEY")
	return token.SignedString([]byte(secretkey))
}

func ValidateToken(w http.ResponseWriter, r *http.Request, tokenString string) (bool, uint) {
	claims := &Claims{}
	secretkey := os.Getenv("JWT_KEY")
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
