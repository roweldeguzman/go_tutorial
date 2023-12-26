package routers

import (
	"api/authorization"
	c "api/controllers/users"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func usersRoute(router *mux.Router) *mux.Router {

	router.Handle("/v1/users/add",
		negroni.New(
			negroni.HandlerFunc(authorization.IsAuthorized),
			negroni.HandlerFunc(c.Create),
		)).
		Methods("POST")

	router.HandleFunc("/v1/users/update", c.Update).Methods("PUT")
	router.HandleFunc("/v1/users/delete", c.Delete).Methods("DELETE")
	router.HandleFunc("/v1/users/get", c.Get).Methods("GET")
	router.HandleFunc("/v1/users/get-info/{id}", c.GetInfo).Methods("GET")

	return router
}
