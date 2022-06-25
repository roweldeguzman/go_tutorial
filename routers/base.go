package routers

import (
	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	router = UsersRoute(router)

	return router
}
