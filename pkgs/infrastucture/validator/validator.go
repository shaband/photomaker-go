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
	ar_uni   *ut.UniversalTranslator
	en_uni   *ut.UniversalTranslator
	// uni      *ut.UniversalTranslator
)

func Validate[T any](data T) map[string]string {
	validate = validator.New()

	en := en.New()
	ar := ar.New()
	en_uni = ut.New(en, en)
	ar_uni = ut.New(ar, ar)
	var trans ut.Translator
	if middleware.CurrentLanguage == "ar" {
		trans, _ = ar_uni.GetTranslator("ar")
		ar_translations.RegisterDefaultTranslations(validate, trans)
	} else {

		trans, _ = en_uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validate, trans)
	}

	err := val(data)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	fmt.Println(errs.Translate(trans))
	return errs.Translate(trans)
}

func val[T any](data T) error {

	err := validate.Struct(data)

	return err
}
