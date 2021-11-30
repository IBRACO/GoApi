package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connexion() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=3308 sslmode=disable "
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("Successfully connected! %T", db)
	return db
}
