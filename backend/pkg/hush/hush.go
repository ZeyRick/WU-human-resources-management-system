package hush

import (
	"golang.org/x/crypto/bcrypt"
)

func Hush(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hushedpassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hushedpassword), []byte(password))
	return err
}
