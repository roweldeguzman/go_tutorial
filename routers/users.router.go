package routers

import (
	"api/authorization"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func usersRoute(router *mux.Router) *mux.Router {

	userRouter := mux.NewRouter().PathPrefix("/v1/users").Subrouter().StrictSlash(true)
	userRouter.HandleFunc("/add", UserController.Create).Methods("POST")
	userRouter.HandleFunc("/update", UserController.Update).Methods("PUT")
	userRouter.HandleFunc("/delete", UserController.Delete).Methods("DELETE")
	userRouter.HandleFunc("/get", UserController.Get).Methods("GET")
	userRouter.HandleFunc("/get-info/{id}", UserController.GetInfo).Methods("GET")
	userRouter.HandleFunc("/search", UserController.SearchUser).Methods("GET")

	router.PathPrefix("/v1/users").Handler(negroni.New(
		negroni.HandlerFunc(authorization.IsAuthorized),
		negroni.Wrap(userRouter),
	))

	return router
}
