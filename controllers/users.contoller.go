package controller

// import (
// 	"api/models"
// 	"api/service"
// 	"api/utils"
// 	"api/validation"
// 	"net/http"
// )

// type controller struct{}

// var userService service.UserService

// // var userService services.UserService

// type IUserController interface {
// 	Create(w http.ResponseWriter, r *http.Request)
// 	Update(w http.ResponseWriter, r *http.Request)
// 	Delete(w http.ResponseWriter, r *http.Request)
// 	// Get(w http.ResponseWriter, r *http.Request)
// 	// GetInfo(w http.ResponseWriter, r *http.Request)
// }

// func NewUserController(service service.UserService) IUserController {
// 	userService = service
// 	return &controller{}
// }

// func (*controller) Create(w http.ResponseWriter, r *http.Request) {
// 	body, mgs := utils.HttpReq(r)
// 	if body == nil {
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": mgs,
// 		}, 500, w)

// 		return
// 	}

// 	firstName, _ := body["firstName"].(string)
// 	lastName, _ := body["lastName"].(string)
// 	email, _ := body["email"].(string)
// 	password, _ := body["password"].(string)
// 	userStatus, _ := body["userStatus"].(string)

// 	user := models.Users{
// 		FirstName:  firstName,
// 		LastName:   lastName,
// 		Email:      email,
// 		Password:   password,
// 		UserStatus: userStatus,
// 	}

// 	validate := validation.Validate()

// 	if err := validate.Struct(user); err != nil {
// 		errs := validation.GetErrors(err)
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": errs,
// 		}, 200, w)
// 		return
// 	}

// 	hashPassword, _ := utils.HashPassword(password)
// 	user.Password = hashPassword

// 	createdUser, err := userService.Create(&user)

// 	if err != nil {
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": err.Error(),
// 		}, 200, w)
// 		return
// 	}

// 	utils.Response(map[string]any{
// 		"statusCode": 200,
// 		"devMessage": createdUser,
// 	}, 200, w)
// }

// func (*controller) Update(w http.ResponseWriter, r *http.Request) {
// 	body, mgs := utils.HttpReq(r)

// 	if body == nil {
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": mgs,
// 		}, 500, w)
// 		return
// 	}

// 	id, _ := body["id"].(float64)
// 	firstName, _ := body["firstName"].(string)
// 	lastName, _ := body["lastName"].(string)
// 	email, _ := body["email"].(string)
// 	userStatus, _ := body["userStatus"].(string)

// 	user := models.Users{
// 		ID:         uint(id),
// 		FirstName:  firstName,
// 		LastName:   lastName,
// 		Email:      email,
// 		UserStatus: userStatus,
// 		Password:   "noValidate",
// 	}

// 	validate := validation.Validate()
// 	err := validate.Struct(user)

// 	if err != nil {
// 		errs := validation.GetErrors(err)
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": errs,
// 		}, 200, w)
// 		return
// 	}

// 	if err := userService.Update(&user); err != nil {
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": err.Error(),
// 		}, 200, w)
// 		return
// 	}

// 	utils.Response(map[string]any{
// 		"statusCode": 200,
// 		"devMessage": user.ID,
// 	}, 200, w)
// }

// func (*controller) Delete(w http.ResponseWriter, r *http.Request) {
// 	body, mgs := utils.HttpReq(r)

// 	if body == nil {
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": mgs,
// 		}, 500, w)
// 		return
// 	}
// 	IDS, _ := body["ids"].([]any)
// 	var ids []uint
// 	for _, id := range IDS {
// 		ID, isValid := id.(float64)
// 		if isValid {
// 			ids = append(ids, uint(ID))
// 		}
// 	}

// 	deletes := models.UserDelete{
// 		IDS: ids,
// 	}
// 	if err := userService.Delete(&deletes); err != nil {
// 		utils.Response(map[string]any{
// 			"statusCode": 500,
// 			"devMessage": err.Error(),
// 		}, 200, w)

// 		return
// 	}
// 	utils.Response(map[string]any{
// 		"statusCode": 200,
// 		"devMessage": "success",
// 	}, 200, w)
// }

// func (*controller) get(w http.ResponseWriter, r *http.Request) {}

// func (*controller) getInfo(w http.ResponseWriter, r *http.Request) {}
