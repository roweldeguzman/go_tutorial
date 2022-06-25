package routers

import (
	u "api/controllers/users"

	"github.com/gorilla/mux"
)

func UsersRoute(router *mux.Router) *mux.Router {

	router.HandleFunc("/v1/users/add", u.Create).Methods("POST")
	router.HandleFunc("/v1/users/update", u.Update).Methods("PUT")
	router.HandleFunc("/v1/users/delete", u.Delete).Methods("DELETE")
	router.HandleFunc("/v1/users/get", u.Get).Methods("GET")
	router.HandleFunc("/v1/users/get-info/{id}", u.GetInfo).Methods("GET")

	return router
}
