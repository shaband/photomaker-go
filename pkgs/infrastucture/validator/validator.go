package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/ar"
	ar_translations "github.com/go-playground/validator/v10/translations/ar"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	ut "github.com/go-playground/universal-translator"
)

var (
	validate *validator.Validate
	// ar_uni   *ut.UniversalTranslator
	// en_uni   *ut.UniversalTranslator
	uni      *ut.UniversalTranslator
)

func Validate[T any](data T) error {
	validate = validator.New()

	en := en.New()
	ar := ar.New()
	// en_uni = ut.New(en, en)
	// ar_uni = ut.New(ar, ar)
	uni = ut.New( en,ar)
	// this is usually know or extracted from http 'Accept-Language' header
	trans, _:= uni.FindTranslator()

	// en_trans, _ := en_uni.GetTranslator("ar")
	// ar_trans, _ := ar_uni.GetTranslator("en")
	validate = validator.New()
	// validate.RegisterTagNameFunc()
	en_translations.RegisterDefaultTranslations(validate, trans)
	ar_translations.RegisterDefaultTranslations(validate, trans)

	err := val(data)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	// fmt.Println(errs.Translate(en_trans))
	fmt.Println(errs.Translate(trans))
	return errs
}

func val[T any](data T) error {

	err := validate.Struct(data)

	return err
}
