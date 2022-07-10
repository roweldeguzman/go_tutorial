package main

import "api/server"

func main() {
	app := server.App{}
	app.Initialize()
	app.Run(":9999")
}
