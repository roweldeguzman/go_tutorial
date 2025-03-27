package users

import (
	"api/models"
	"api/utils"
	"net/http"
)

func (c *UserController) FindUser(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("orderBy") // desc or asc

	user := models.Users{
		Email: query,
	}
	users, err := c.service.FindUser(&user)

	if err != nil {
		// if err := user.GetInfo(); err != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]any{
		"statusCode": 200,
		"devMessage": users,
	}, 200, w)
}
