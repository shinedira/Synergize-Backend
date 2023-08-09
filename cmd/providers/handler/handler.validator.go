package handler

import (
	"fmt"
	"net/http"
	"synergize/backend-test/cmd/providers/handler/validations"
	"synergize/backend-test/pkg/helper"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	var messageBag []string

	if err := cv.validator.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("The %s field is %s", e.Field(), e.ActualTag())
			messageBag = append(messageBag, message)
		}

		return helper.ErrorReponse(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "Unprocessable Entity",
			"errors":  messageBag,
		})
	}

	return nil
}

func bootCustomeValidation(v *validator.Validate) {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	validations.BootDBValidation(v, trans)
}
