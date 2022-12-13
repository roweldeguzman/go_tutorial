package users

import (
	"api/models"
	"api/utils"
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	page := utils.PagerTernary(r.FormValue("page"), 1)
	rows := utils.PagerTernary(r.FormValue("rows"), 10)

	TblUsers := models.TblUsers{}

	users, total, err := TblUsers.Get(r)

	if err != nil {

		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)

		return
	}

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": users,
		"paginate":   utils.Paginate(rows, page, int(total)),
	}, utils.Code.OK, w)

}
