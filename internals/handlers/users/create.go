package users

import (
	_ "api/internals/util/context"
	"api/models"
	"api/utils"
	"api/validation"
	"net/http"
)

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {

	// reqID := ctx.RequestID(r.Context())

	body, mgs := utils.HttpReq(r)
	if body == nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": mgs,
		}, 500, w)

		return
	}

	firstName, _ := body["firstName"].(string)
	lastName, _ := body["lastName"].(string)
	email, _ := body["email"].(string)
	password, _ := body["password"].(string)
	userStatus, _ := body["userStatus"].(string)

	user := models.Users{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Password:   password,
		UserStatus: userStatus,
	}

	validate := validation.Validate()

	if err := validate.Struct(user); err != nil {
		errs := validation.GetErrors(err)
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": errs,
		}, 200, w)
		return
	}

	hashPassword, _ := utils.HashPassword(password)
	user.Password = hashPassword

	userId, err := c.service.Create(&user)

	if err != nil {
		utils.Response(map[string]any{
			"statusCode": 500,
			"devMessage": err.Error(),
		}, 200, w)
		return
	}

	utils.Response(map[string]any{
		"statusCode": 200,
		"devMessage": userId,
	}, 200, w)
}
