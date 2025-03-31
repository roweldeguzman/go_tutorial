package models

import (
	"api/database"
	"database/sql"
	"time"
)

var DB = database.DB

type DateModel struct {
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`
}
