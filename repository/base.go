package repository

import (
	"api/database"
	"api/struct/pagination"
	"slices"

	"gorm.io/gorm"
)

var DB = database.DB

func paginate(pageParam pagination.PagingOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pageParam.Page
		rows := pageParam.Rows

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

func order(sortParam pagination.SortingOptions, fields []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		orderBy := sortParam.OrderBy // desc or asc
		sortBy := sortParam.SortBy   // by model fields
		groupBy := sortParam.GroupBy // by model fields

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
