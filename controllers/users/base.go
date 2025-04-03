package users

import (
	"api/service"
)

type UserController struct {
	service *service.UsersService
}

func NewUserController(service *service.UsersService) *UserController {

	return &UserController{service}
}
