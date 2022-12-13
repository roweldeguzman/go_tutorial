package server

import (
	"api/database"
	"api/models"
	"api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

type ResponseCode struct {
	OK       int
	EXIST    int
	NOTFOUND int
	ISE      int
	INVALID  int
}

type App struct {
	Router *mux.Router
}

var responseCode = ResponseCode{
	OK:       200,
	EXIST:    201,
	NOTFOUND: 404,
	ISE:      500,
	INVALID:  406,
}

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}
}

func (a *App) Initialize() {
	c := cors.New(cors.Options{
		AllowCredentials:   true,
		OptionsPassthrough: false,
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedOrigins:     []string{"*"},
	})

	initDatabase()

	models.DB = database.Connect
	a.Router = routers.LoadRouter()

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(a.Router)

}

func (a *App) Run(port string) {
	fmt.Print("http://localhost" + port + "\n")
	log.Fatal(http.ListenAndServe(port, a.Router))
}
