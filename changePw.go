package main

import (
	// "fmt"
	"net/http"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type data struct {
	Userdb  *User
	Msg    string
}
type UserChangePassword struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func ChangePassword(c echo.Context) error {
	user := new(User)
	userFront := new(UserChangePassword)
	if err := c.Bind(userFront); err != nil {
		return err
	}
	result := db.Where("email = ?", userFront.Email).First(&user)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "bad email adresse")
	}
	match, err := ComparePassword(userFront.Password, user.Password)
	if !match || err != nil {
		return c.String(http.StatusBadRequest, "bad password")
	}

	hash, err := GeneratePassword(userFront.NewPassword)
	if err != nil {
		panic(err)
	}

	user.Password = hash
	db.Save(&user)

	
	return c.JSON(http.StatusOK,data{user,"Vous avez bien changer votre mot de pass"+user.First_name})
}
