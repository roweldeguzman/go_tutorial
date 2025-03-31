package users

import (
	"api/struct/pagination"
	"api/struct/response"
	"api/utils"
	"net/http"
)

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	page := utils.Ternary(r.URL.Query().Get("page"), 1)
	rows := utils.Ternary(r.URL.Query().Get("rows"), 10)
	orderBy := r.URL.Query().Get("orderBy") // desc or asc
	sortBy := r.URL.Query().Get("sortBy")   // by model fields
	groupBy := r.URL.Query().Get("groupBy") // by model fields

	pageParams := pagination.PagingOptions{
		Page: page,
		Rows: rows,
	}

	sortParams := pagination.SortingOptions{
		OrderBy: orderBy,
		SortBy:  sortBy,
		GroupBy: groupBy,
	}

	users, total, err := c.service.Get(pageParams, sortParams)

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
		"paginate":   utils.Paginate(rows, page, int(total)),
	}, response.Code.OK, w)
}
