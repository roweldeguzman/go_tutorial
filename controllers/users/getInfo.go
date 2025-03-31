package users

import (
	"api/models"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *UserController) GetInfo(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID, isValidInt := strconv.Atoi(param["id"])

	if isValidInt != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": "Unable to find user.",
		}, 200, w)
		return
	}

	user := models.Users{
		ID: uint(ID),
	}

	if err := c.service.GetInfo(&user); err != nil {
		// if err := user.GetInfo(); err != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]any{
		"statusCode": 200,
		"devMessage": user,
	}, 200, w)
}
