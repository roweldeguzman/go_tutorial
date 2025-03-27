package main

import (
	"api/database"
	"api/models"
	"fmt"
)

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}
}

func main() {
	initDatabase()

	var db = database.DB
	models.DB = database.DB

	fmt.Println("Connected! Starting migration...")

	defer database.Close()
	TblUsers(db)

	fmt.Println("Migration Finished...")
}
