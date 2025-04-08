package main

import (
	"api/database"
	"api/models"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}
}

func main() {
	initDatabase()

	db = database.DB
	models.DB = database.DB

	fmt.Println("Connected! Starting migration...")

	defer database.Close()
	createTable()

	// tblUsers()

	fmt.Println("Migration Finished...")
}
