package main

import (
	"api/database"
	"api/server"
	"fmt"
)

func main() {
	app := server.App{}

	defer cleanUp()

	db, err := database.NewDatabaseConnection()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	app.Initialize(db.DB)
	app.Run(":9999")
}

func cleanUp() {
	if r := recover(); r != nil {
		fmt.Println(r) // Replace with logging system
	}
}
