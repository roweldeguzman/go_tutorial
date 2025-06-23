package controller

import (
	"api/authorization"
	"api/models"
	"api/service"
	"api/utils"
	"fmt"
	"net/http"
	"strings"
)

type AuthController struct {
	service *service.UsersService
}

func NewAuthController(service *service.UsersService) *AuthController {

	return &AuthController{service}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	body, mgs := utils.HttpReq(r)

	if body == nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": "Invalid username or password. " + mgs,
		}, 500, w)
		return
	}
	email, _ := body["email"].(string)
	password, _ := body["password"].(string)

	if strings.Trim(email, " ") == "" || strings.Trim(password, "") == "" {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": "Wrong username or password.",
		}, 200, w)
		return
	}
	user := models.Users{Email: email}

	if user, err := c.service.FindUser(&user); err != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	} else if err := utils.ComparePasswords(user.Password, password); err != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": "Wrong username or Password.",
		}, 200, w)
		return
	}

	fmt.Println(user)
	token, err := authorization.GenerateJWT(map[string]any{
		"email": user.Email,
		"id":    user.ID,
	})

	if err == nil {
		utils.Response(map[string]any{
			"statusCode": 200,
			"devMessage": map[string]any{
				"token": token,
				"user":  user,
			},
		}, 200, w)
		return
	}

	utils.Response(map[string]any{
		"statusCode": 500,
		"devMessage": err,
	}, 200, w)

}
