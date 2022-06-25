package main

import (
	"api/models"
	"api/utils"

	"gorm.io/gorm"
)

func TblUsers(db *gorm.DB) {

	table := &models.TblUsers{}
	columns := models.TblUsers{
		FirstName:  "Rowel",
		LastName:   "de Guzman",
		Email:      "rowel.deguzman@roweldev.com",
		Password:   utils.MakePassword("12345"),
		UserStatus: "1",
	}
	if exist := db.Migrator().HasTable(table); !exist {
		utils.ErrorChecker(0, db.AutoMigrate(table))

		utils.ErrorChecker(0, columns.Create())
	}

}
