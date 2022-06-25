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
	return validate
}

func GetErrors(err error) []string {
	var errorList []string
	trans, _ := uni.GetTranslator("en")

	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errorList = append(errorList, e.Translate(trans))
	}

	return errorList
}
