package server

import (
	"api/database"
	"api/models"
	"api/repository"
	"api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

type App struct {
	Router *mux.Router
}

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}
}

func (app *App) Initialize() {

	c := cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedOrigins:     []string{"*"},
	})

	initDatabase()

	models.DB = database.DB
	repository.DB = database.DB

	app.Router = routers.LoadRouter()

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(app.Router)

}

func (app *App) Run(port string) {

	fmt.Print("http://localhost" + port + "\n")
	log.Fatal(http.ListenAndServe(port, app.Router))
}
