package infrastructure

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return string(hashedPassword), nil
}

func ComparePassword(correctPassword []byte, inputPassword []byte) error {
	if bcrypt.CompareHashAndPassword(correctPassword, inputPassword) != nil {
		return errors.New("invalid username or password")
	}
	return nil
}