package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/ar"
	ar_translations "github.com/go-playground/validator/v10/translations/ar"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"

	ut "github.com/go-playground/universal-translator"
)

var (
	validate *validator.Validate
	// ar_uni   *ut.UniversalTranslator
	// en_uni   *ut.UniversalTranslator
	uni *ut.UniversalTranslator
)

func Validate[T any](data T) error {
	validate = validator.New()

	en := en.New()
	ar := ar.New()
	uni = ut.New(en, ar)
	// ar_uni = ut.New(ar, ar)
	// var trans ut.Translator
	trans, _ := uni.GetTranslator(middleware.CurrentLanguage)
	if middleware.CurrentLanguage == "ar" {
		ar_translations.RegisterDefaultTranslations(validate, trans)
	} else {
		en_translations.RegisterDefaultTranslations(validate, trans)
	}

	err := val(data)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	fmt.Println(errs.Translate(trans))
	return errs
}

func val[T any](data T) error {

	err := validate.Struct(data)

	return err
}
