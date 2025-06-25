package validator

import (
	"fmt"

	"github.com/featriadi/golang-libs/utils"
	"github.com/go-playground/validator/v10"
)

func Validator(structTarget any) map[string]string {
	validate := validator.New()
	validate.RegisterValidation("date", dateValidator)

	if err := validate.Struct(structTarget); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := utils.GetJsonTagOnField(structTarget, err.StructField())

			errors[fieldName] = fmt.Sprintf("Field %s failed validation on %s condition", fieldName, err.Tag())
		}

		return errors
	}

	return nil
}

func dateValidator(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}

	dateString := fl.Field().String()
	time := utils.ParseStringToDate(dateString)

	return time != nil
}
