package server

import (
	"api/models"
	"api/repository"
	"api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize(db *gorm.DB) {

	c := cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedOrigins:     []string{"*"},
	})

	models.DB = db
	repository.DB = db

	app.Router = routers.LoadRouter()

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(app.Router)

}

func (app *App) Run(port string) {

	fmt.Print("http://localhost" + port + "\n")
	log.Fatal(http.ListenAndServe(port, app.Router))
}
