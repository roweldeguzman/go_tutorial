package routers

import (
	controller "api/controllers"
	"api/repository"
	"api/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Router struct {
	mux            *mux.Router
	authController *controller.AuthController
}

func NewRouter(db *gorm.DB) *Router {
	usersRepository := repository.NewUsersRepository(db)
	usersService := service.NewUsersService(usersRepository)
	authController := controller.NewAuthController(usersService)

	return &Router{
		mux:            mux.NewRouter(),
		authController: authController,
	}
}

func (r *Router) SetupRouter() *mux.Router {
	r.mux.HandleFunc("/v1/auth/login", r.authController.Login).Methods("POST")

	return r.mux
}

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
