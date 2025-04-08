package main

import (
	"api/models"
	"api/utils"
)

func createTable() {
	var tableList = map[string]any{
		// "users": &models.Users{},
		"posts": &models.Posts{},
	}

	for _, table := range tableList {

		//exist := db.Migrator().HasTable(table)

		utils.ErrorChecker(0, db.AutoMigrate(table))

	}
}
