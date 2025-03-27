package routers

import (
	"api/authorization"
	"api/controllers/users"
	users2 "api/controllers/users2"
	"api/repository"
	"api/service"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	userRepository *repository.UserRepository = repository.NewUserRepository()
	userService    *service.UserService       = service.NewUserService(userRepository)
	UserController *users2.UserController     = users2.NewUserController(userService)
)

func usersRoute(router *mux.Router) *mux.Router {

	userRouter := mux.NewRouter().PathPrefix("/v1/users").Subrouter().StrictSlash(true)
	userRouter.HandleFunc("/add", UserController.Create).Methods("POST")
	userRouter.HandleFunc("/update", UserController.Update).Methods("PUT")
	userRouter.HandleFunc("/delete", UserController.Delete).Methods("DELETE")
	userRouter.HandleFunc("/get", UserController.Get).Methods("GET")
	userRouter.HandleFunc("/get-info/{id}", users.GetInfo).Methods("GET")

	router.PathPrefix("/v1/users").Handler(negroni.New(
		negroni.HandlerFunc(authorization.IsAuthorized),
		negroni.Wrap(userRouter),
	))

	// router.Handle("/v1/users/add",
	// 	negroni.New(
	// 		negroni.HandlerFunc(authorization.IsAuthorized),
	// 		negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// 			c.Create(w, r)
	// 		}),
	// 	)).
	// 	Methods("POST")

	// router.HandleFunc("/v1/users/update", c.Update).Methods("PUT")
	// router.HandleFunc("/v1/users/delete", c.Delete).Methods("DELETE")
	// router.HandleFunc("/v1/users/get", c.Get).Methods("GET")
	// router.HandleFunc("/v1/users/get-info/{id}", c.GetInfo).Methods("GET")

	return router
}
