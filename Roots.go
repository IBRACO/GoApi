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
	e.POST("/uppass", UpdatePassword)
	e.POST("/save", SaveUser)
	e.POST("/login", Login)
	e.Logger.Fatal(e.Start(":1323"))
}


//2bEg7x8Xe/lGUod5oABugg$kAtivV1zuxfl7KcsT7EnClxRPGSbzMN2TkncxiEQS1w