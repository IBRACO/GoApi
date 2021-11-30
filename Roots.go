package main

import "github.com/labstack/echo/v4"


func Roots() {
	db = connexion()
	

	e := echo.New()
	e.GET("/user/:id", GetUserById)
	e.PUT("/user/update/:id", UpdateUserById)
	e.DELETE("/user/delete/:id", DeleteUserById)
	e.POST("/fg", ForgetPw)
	e.POST("/change", ChangePassword)
	e.POST("/save", SaveUser)
	e.POST("/login", Login)
	e.Logger.Fatal(e.Start(":1323"))
}


