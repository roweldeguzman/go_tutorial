package models

import (
	"api/database"
	"api/utils"
	"database/sql"
	"net/http"
	"time"

	"gorm.io/gorm"
)

var DB = database.Connect

type DateModel struct {
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}

func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := utils.PagerTernary(r.FormValue("page"), 1)
		rows := utils.PagerTernary(r.FormValue("rows"), 10)

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

		if orderBy == "" || !utils.InArray(orderBy, []string{"desc", "asc"}) {
			orderBy = "desc"
		}
		if sortBy == "" || !utils.InArray(sortBy, fields) {
			sortBy = "id"
		}

		ctx := db.Order(sortBy + " " + orderBy)

		if groupBy != "" && utils.InArray(orderBy, []string{"name"}) {
			ctx = ctx.Group(groupBy)
		}

		return ctx
	}
}
