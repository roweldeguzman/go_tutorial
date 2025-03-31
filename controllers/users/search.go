package users

import (
	"api/models"
	"api/utils"
	"fmt"
	"net/http"
)

func (c *UserController) SearchUser(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query") // desc or asc

	fmt.Println(query)
	user := models.Users{
		Email: query,
	}
	users, err := c.service.SearchUser(&user)

	if err != nil {
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
