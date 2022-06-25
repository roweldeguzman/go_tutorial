package users

import (
	"api/models"
	"api/utils"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	body, mgs := utils.HttpReq(r)

	if body == nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": mgs,
		}, 500, w)
		return
	}
	IDS, _ := body["ids"].([]interface{})
	var ids []uint
	for _, id := range IDS {
		ID, isValid := id.(float64)
		if isValid {
			ids = append(ids, uint(ID))
		}
	}

	deletes := models.TblUserDelete{
		IDS: ids,
	}
	if err := deletes.Delete(); err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)

		return
	}
	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": "success",
	}, 200, w)

}
