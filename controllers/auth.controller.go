package controllers

import (
	"api/models"
	"api/security"
	"api/utils"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, mgs := utils.HttpReq(r)

	if body == nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "Invalid username or password " + mgs,
		}, 500, w)
		return
	}
	email, _ := body["email"].(string)
	password, _ := body["password"].(string)

	if strings.Trim(email, " ") == "" || strings.Trim(password, "") == "" {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "Wrong username or password.",
		}, 200, w)
		return
	}
	user := models.TblUsers{
		Email:    email,
		Password: password,
	}

	if err := user.FindUser(); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	token, err := security.GenerateJWT(map[string]interface{}{
		"email": user.Email,
		"id":    user.ID,
	})

	if err == nil {
		utils.Response(map[string]interface{}{
			"statusCode": 200,
			"devMessage": map[string]interface{}{
				"token": token,
				"user":  user,
			},
		}, 200, w)
		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 500,
		"devMessage": err,
	}, 200, w)

}
