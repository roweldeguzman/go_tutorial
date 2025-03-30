package main

import (
	"api/models"
	"api/utils"
)

func createTable() {
	var tableList = map[string]any{
		"user": &models.Users{},
	}

	for _, table := range tableList {

		//exist := db.Migrator().HasTable(table)

		utils.ErrorChecker(0, db.AutoMigrate(table))

	}
}
