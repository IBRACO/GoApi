package main

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "os/exec"
	"strings"
	// "strconv"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func ForgetPw(c echo.Context) error {
	// db := connexion()
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	result1 := db.Where("email = ?", user.Email).First(&user)
	if result1.Error != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%v", result1.Error))
	}

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	user.UUID = uuid
	// user.Token = token
	db.Save(&user)

	//send mail with token
	sendmail(user.Email, user.UUID)

	return c.JSON(http.StatusOK, user)
}

func UpdatePassword(c echo.Context) error {
	// db := connexion()

	user := new(User)
	// user2 := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	//  token := c.FormValue("token")
	passwordSended := user.Password

	result := db.Where("uuid = ?", user.UUID).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", result.Error))
	}

	hash,err := GeneratePassword(passwordSended)
	if err != nil {
		panic(err)
	}


	user.Password =hash
	db.Save(&user)
	sendmail(user.Email, "vous avez changez votre mot de pass avec succee")

	return c.String(http.StatusOK, user.Password)
}
