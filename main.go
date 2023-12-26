package main

import (
	"api/database"
	"api/server"
)

func main() {
	app := server.App{}
	app.Initialize()

	defer database.Close()

	app.Run(":9999")
}
