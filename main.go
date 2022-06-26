package main

import (
	"api/database"
	"api/models"
	"api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}

	fmt.Println("Database connection successfully opened.")
}

func main() {

	c := cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedOrigins:     []string{"*"},
	})
	router := routers.LoadRouter()

	initDatabase()

	defer database.Close()

	models.DB = database.Connect

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	fmt.Print("http://localhost:9999\n")
	log.Fatal(http.ListenAndServe(":9999", n))
}
