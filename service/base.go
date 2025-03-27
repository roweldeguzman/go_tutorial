package service

import (
	"api/utils"
	"net/http"
	"slices"

	"gorm.io/gorm"
)

func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := utils.Ternary(r.FormValue("page"), 1)
		rows := utils.Ternary(r.FormValue("rows"), 10)

		if page == 0 {
			page = 1
		}

		switch {
		case rows > 100:
			rows = 100
		case rows <= 0:
			rows = 10
		}

		offset := (page - 1) * rows
		return db.Offset(offset).Limit(rows)
	}
}

func order(r *http.Request, fields []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		orderBy := r.URL.Query().Get("orderBy") // desc or asc
		sortBy := r.URL.Query().Get("sortBy")   // by model fields
		groupBy := r.URL.Query().Get("groupBy") // by model fields

		if orderBy == "" || !slices.Contains([]string{"desc", "asc"}, orderBy) {
			orderBy = "desc"
		}
		if sortBy == "" || !slices.Contains(fields, sortBy) {
			sortBy = "id"
		}

		ctx := db.Order(sortBy + " " + orderBy)

		if groupBy != "" && slices.Contains([]string{"name"}, orderBy) {
			ctx = ctx.Group(groupBy)
		}

		return ctx
	}
}
