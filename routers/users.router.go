package routers

import (
	"api/authorization"
	controller "api/controllers/users"
	"api/repository"
	"api/service"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	userRepository = repository.NewUserRepository()
	userService    = service.NewUserService(userRepository)
	UserController = controller.NewUserController(userService)
)

func usersRoute(router *mux.Router) *mux.Router {

	userRouter := mux.NewRouter().PathPrefix("/v1/users").Subrouter().StrictSlash(true)
	userRouter.HandleFunc("/add", UserController.Create).Methods("POST")
	userRouter.HandleFunc("/update", UserController.Update).Methods("PUT")
	userRouter.HandleFunc("/delete", UserController.Delete).Methods("DELETE")
	userRouter.HandleFunc("/get", UserController.Get).Methods("GET")
	// userRouter.HandleFunc("/get-info/{id}", users.GetInfo).Methods("GET")

	router.PathPrefix("/v1/users").Handler(negroni.New(
		negroni.HandlerFunc(authorization.IsAuthorized),
		negroni.Wrap(userRouter),
	))

	return router
}
