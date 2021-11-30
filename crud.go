package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserById(c echo.Context) error {
	db := connexion()
	id := c.Param("id")
	user := User{}
	db.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

func UpdateUserById(c echo.Context) error {
	db := connexion()
	// User ID from path `users/:id`
	id := c.Param("id")
	user := User{}
	db.First(&user, id)
	//Form data not json
	user.Email = c.FormValue("email")
	user.Age = c.FormValue("age")
	db.Save(&user)
	return c.String(http.StatusOK, user.First_name)
}

func DeleteUserById(c echo.Context) error {
	db := connexion()
	id := c.Param("id")
	st := db.Delete(&User{}, id)
	if int(st.RowsAffected) == 0 {
		return c.String(http.StatusBadRequest, "The user ID not fund")
	}

	return c.String(http.StatusOK, fmt.Sprintf("The user with ID %s has been deleted successfully\n  %d line Deleted", id, st.RowsAffected))
}
