package routers

import (
	"api/authorization"
	u "api/controllers/users"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func usersRoute(router *mux.Router) *mux.Router {

	router.Handle("/v1/users/add",
		negroni.New(
			negroni.HandlerFunc(authorization.IsAuthorized),
			negroni.HandlerFunc(u.Create),
		)).
		Methods("POST")

	router.HandleFunc("/v1/users/update", u.Update).Methods("PUT")
	router.HandleFunc("/v1/users/delete", u.Delete).Methods("DELETE")
	router.HandleFunc("/v1/users/get", u.Get).Methods("GET")
	router.HandleFunc("/v1/users/get-info/{id}", u.GetInfo).Methods("GET")

	return router
}
