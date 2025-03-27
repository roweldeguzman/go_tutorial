package users

import (
	"api/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {

	return &UserController{service}
}
