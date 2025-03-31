package routers

import (
	controller "api/controllers/users"
	"api/repository"
	"api/service"

	"github.com/gorilla/mux"
)

var (
	userRepository = repository.NewUserRepository()
	userService    = service.NewUserService(userRepository)
	UserController = controller.NewUserController(userService)
)

func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	router = usersRoute(router)
	router = auth(router)

	return router
}
