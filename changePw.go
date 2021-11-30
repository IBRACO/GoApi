package main

import (
	"fmt"
	"net/http"

	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ChangePassword(c echo.Context) error {
	db := connexion()

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

// func Verify(c echo.Context) error {
// 	db := connexion()
// 	token := c.FormValue("token")

// 	tokenString := token
// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("AllYourBase"), nil
// 	})
// 	// ... error handling
// 	if err != nil {
// 		return err
// 	}

// 	m := make(map[string]interface{})
// 	// do something with decoded claims
// 	for key, val := range claims {
// 		m[key] = val
// 	}
// 	// id := m["id"].(string)
// 	user := new(User)
// 	result1 := db.First(&user, 9)
// 	if result1.Error != nil {
// 		return c.String(http.StatusBadRequest, fmt.Sprintf("%v", result1.Error))
// 	}

// 	sendmail(user.Email,"")
// 	// return c.JSON(http.StatusOK, m)
// 	return c.String(http.StatusOK, "You receved a verification code")
// }
