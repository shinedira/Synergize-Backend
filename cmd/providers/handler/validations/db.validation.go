package validations

import (
	"strings"
	"synergize/backend-test/pkg/facades"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func BootDBValidation(v *validator.Validate, trans ut.Translator) {
	v.RegisterValidation("unique", ValidateUnique)

	v.RegisterTranslation("unique", trans, func(ut ut.Translator) error {
		return ut.Add("unique", "{0} already exists", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique", fe.Field())

		return t
	})
}
func ValidateUnique(fl validator.FieldLevel) bool {
	tagParts := strings.Split(fl.Param(), ":")

	if len(tagParts) != 2 {
		return false
	}

	tableName := tagParts[0]
	fieldName := tagParts[1]
	fieldValue := fl.Field().String()

	var count int64
	facades.DB.Table(tableName).Where(fieldName+" = ?", fieldValue).Count(&count)

	return count == 0
}
