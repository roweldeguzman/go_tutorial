package validation

import (
	"api/utils"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate() *validator.Validate {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()

	utils.ErrorChecker(0, en_translations.RegisterDefaultTranslations(validate, trans))
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("error"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	utils.ErrorChecker(0, validate.RegisterValidation("validPassword", func(fld validator.FieldLevel) bool {

		value := fld.Field().String()
		// check if using user struct for update
		if value == "noValidate" {
			return true
		}
		if err := utils.ValidatePassword(value); err != nil {
			return false

		}

		return true
	}))

	return validate
}

func GetErrors(err error) []string {
	var errorList []string
	trans, _ := uni.GetTranslator("en")

	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		if e.Field() == "Password" {
			errorList = append(errorList, "Password did not meet the required strength. Must be 8-32 characters long, must contain 1 letter, must contain 1 number")
		} else {
			errorList = append(errorList, e.Translate(trans))
		}
	}

	return errorList
}
