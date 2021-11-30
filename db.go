package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

 var db *gorm.DB



func connexion() *gorm.DB{
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=3308 sslmode=disable "
	ds, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("Successfully connected! %T", ds)
	return ds
	

}