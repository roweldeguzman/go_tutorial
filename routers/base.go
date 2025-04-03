package routers

import (
	controller "api/controllers/users"
	"api/repository"
	"api/service"

	"github.com/gorilla/mux"
)

var (
	usersRepository = repository.NewUsersRepository()
	usersService    = service.NewUsersService(usersRepository)
	UsersController = controller.NewUserController(usersService)
)

func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	router = usersRoute(router)
	router = auth(router)

	return router
}
