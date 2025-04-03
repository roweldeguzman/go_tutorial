package routers

import (
	controller "api/controllers"

	"github.com/gorilla/mux"
)

var AuthController = controller.NewAuthController(usersService)

func auth(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/auth/login", AuthController.Login).Methods("POST")
	// router.HandleFunc("/v1/auth/test", UserController.Create).Methods("POST")
	return router
}
