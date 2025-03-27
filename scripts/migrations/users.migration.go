package main

import (
	"api/models"
	"api/utils"

	"gorm.io/gorm"
)

func TblUsers(db *gorm.DB) {

	hashPassword, _ := utils.HashPassword("admin")
	table := &models.Users{}
	columns := models.Users{
		FirstName:  "Rowel",
		LastName:   "de Guzman",
		Email:      "rowel.deguzman@roweldev.com",
		Password:   hashPassword,
		UserStatus: "1",
	}
	if exist := db.Migrator().HasTable(table); !exist {
		utils.ErrorChecker(0, db.AutoMigrate(table))

		utils.ErrorChecker(0, columns.Create())
	}

}
