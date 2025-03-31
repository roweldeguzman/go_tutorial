package main

import (
	"api/models"
	"api/utils"
	"fmt"
)

func tblUsers() {

	hashPassword, _ := utils.HashPassword("admin")
	columns := models.Users{
		FirstName:  "Rowel",
		LastName:   "de Guzman",
		Email:      "rowel.deguzman@roweldev.com",
		Password:   hashPassword,
		UserStatus: "1",
	}

	ctx := db.Create(&columns)

	if ctx.Error != nil {
		fmt.Println(ctx.Error)
	}
}
