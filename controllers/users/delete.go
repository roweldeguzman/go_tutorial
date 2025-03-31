package users

import (
	"api/models"
	"api/utils"
	"net/http"
)

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	body, mgs := utils.HttpReq(r)

	if body == nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": mgs,
		}, 500, w)
		return
	}
	IDS, _ := body["ids"].([]any)
	var ids []uint
	for _, id := range IDS {
		ID, isValid := id.(float64)
		if isValid {
			ids = append(ids, uint(ID))
		}
	}

	deletes := models.UserDelete{
		IDS: ids,
	}
	if err := c.service.Delete(&deletes); err != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)

		return
	}
	utils.Response(map[string]any{
		"statusCode": 200,
		"devMessage": "success",
	}, 200, w)
}
