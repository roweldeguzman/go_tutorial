package controller

import (
	"api/authorization"
	"api/models"
	"api/utils"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, mgs := utils.HttpReq(r)

	if body == nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "Invalid username or password. " + mgs,
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
	user := models.Users{Email: email}

	if err := user.FindUser(); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	} else if err := utils.ComparePasswords(user.Password, password); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": "Wrong username or Password.",
		}, 200, w)
		return
	}

	token, err := authorization.GenerateJWT(map[string]interface{}{
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
