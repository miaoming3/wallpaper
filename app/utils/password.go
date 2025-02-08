package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func ComparePoserPassword(hash_password string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	return err == nil
}
