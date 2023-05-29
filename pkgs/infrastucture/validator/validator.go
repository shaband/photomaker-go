package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/ar"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	ar_translations "github.com/go-playground/validator/v10/translations/ar"

	ut "github.com/go-playground/universal-translator"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
)

func Validate[T any](data T) error {
	validate = validator.New()

	en := en.New()
	ar := ar.New()
	uni = ut.New(en, en, ar)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	en_trans, _ := uni.GetTranslator("ar")
	ar_trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	ar_translations.RegisterDefaultTranslations(validate, ar_trans)
	en_translations.RegisterDefaultTranslations(validate, en_trans)

	return val(data)
}

func val[T any](data T) error {

	err := validate.Struct(data)
	return err
}
