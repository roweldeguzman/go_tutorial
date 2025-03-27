package main

import (
	"api/server"
	"fmt"
)

func main() {
	app := server.App{}

	defer cleanUp()

	app.Initialize()
	app.Run(":9999")
}

func cleanUp() {
	if r := recover(); r != nil {
		fmt.Println(r) // Replace with logging system
	}
}
