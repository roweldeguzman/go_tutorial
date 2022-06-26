package routers

import (
	c "api/controllers"

	"github.com/gorilla/mux"
)

func auth(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/auth/login", c.Login).Methods("POST")
	return router
}
