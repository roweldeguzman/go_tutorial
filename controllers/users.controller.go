package controller

// import (
// 	"api/models"
// 	"api/utils"
// 	"api/validation"
// 	"fmt"
// 	"net/http"
// )

// type Users struct{}

// func (c *Users) Create(w http.ResponseWriter, r *http.Request) {
// 	body, mgs := utils.HttpReq(r)

// 	if body == nil {
// 		utils.Response(map[string]interface{}{
// 			"statusCode": 500,
// 			"devMessage": mgs,
// 		}, 500, w)
// 		return
// 	}

// 	firstName, _ := body["firstName"].(string)
// 	lastName, _ := body["lastName"].(string)
// 	email, _ := body["email"].(string)
// 	password, _ := body["password"].(string)
// 	userStatus, _ := body["userStatus"].(float64)
// 	status := fmt.Sprintf("%f", userStatus)

// 	user := models.TblUsers{
// 		FirstName:  firstName,
// 		LastName:   lastName,
// 		Email:      email,
// 		Password:   password,
// 		UserStatus: status,
// 	}
// 	validate := validation.Validate()
// 	err := validate.Struct(user)

// 	if err != nil {
// 		errs := validation.GetErrors(err)
// 		utils.Response(map[string]interface{}{
// 			"statusCode": 500,
// 			"devMessage": errs,
// 		}, 200, w)
// 		return
// 	}

// 	if err := user.Create(); err != nil {
// 		utils.Response(map[string]interface{}{
// 			"statusCode": 500,
// 			"devMessage": err.Error(),
// 		}, 200, w)
// 		return
// 	}

// 	utils.Response(map[string]interface{}{
// 		"statusCode": 200,
// 		"devMessage": user.ID,
// 	}, 200, w)

// }
