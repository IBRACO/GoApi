package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("AllYourBase")

func GenerateJWT(id int) (string, error) {

	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token, err := t.SignedString(SECRET_KEY)
	return token, err
}