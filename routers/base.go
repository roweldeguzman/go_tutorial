package routers

import (
	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	router = usersRoute(router)
	router = auth(router)

	return router
}
