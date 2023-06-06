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
	ar_uni   *ut.UniversalTranslator
	en_uni   *ut.UniversalTranslator
)

func Validate[T any](data T) error {
	validate = validator.New()

	en := en.New()
	ar := ar.New()
	en_uni = ut.New(en, en)
	ar_uni = ut.New(ar, ar)
	// this is usually know or extracted from http 'Accept-Language' header

	en_trans, _ := en_uni.GetTranslator("ar")
	ar_trans, _ := ar_uni.GetTranslator("en")

	validate = validator.New()
	ar_translations.RegisterDefaultTranslations(validate, ar_trans)
	en_translations.RegisterDefaultTranslations(validate, en_trans)

	return val(data)
}

func val[T any](data T) error {

	err := validate.Struct(data)
	return err
}

func enTranslationFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(fe.Tag(), fe.Field())
	return t
}

func arTranslationFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T(fe.Tag(), fe.Field())
	return t
}
