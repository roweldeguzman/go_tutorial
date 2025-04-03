package routers

import (
	"api/authorization"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func usersRoute(router *mux.Router) *mux.Router {

	userRouter := mux.NewRouter().PathPrefix("/v1/users").Subrouter().StrictSlash(true)
	userRouter.HandleFunc("/add", UsersController.Create).Methods("POST")
	userRouter.HandleFunc("/update", UsersController.Update).Methods("PUT")
	userRouter.HandleFunc("/delete", UsersController.Delete).Methods("DELETE")
	userRouter.HandleFunc("/get", UsersController.Get).Methods("GET")
	userRouter.HandleFunc("/get-info/{id}", UsersController.GetInfo).Methods("GET")
	userRouter.HandleFunc("/search", UsersController.SearchUser).Methods("GET")

	router.PathPrefix("/v1/users").Handler(negroni.New(
		negroni.HandlerFunc(authorization.IsAuthorized),
		negroni.Wrap(userRouter),
	))

	return router
}
