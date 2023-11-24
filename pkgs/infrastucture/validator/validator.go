package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/ar"
	ar_translations "github.com/go-playground/validator/v10/translations/ar"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	// "github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"

	ut "github.com/go-playground/universal-translator"
)

var (
	validate *validator.Validate
	ar_uni   *ut.UniversalTranslator
	en_uni   *ut.UniversalTranslator
)

// type myDBAbstraction struct {
// 	db string
// }

//	func (a *myDBAbstraction) ValidateUser(fl validator.FieldLevel) bool {
//		return fl.Field().String() == a.db
//	}
func Validate[T any](data T) map[string]string {
	validate = validator.New()
	// validate.RegisterValidation("is-awesome", func(fl validator.FieldLevel) bool {
	// 	return fl.Field().String() == myDBHandle
	// })
	en := en.New()
	ar := ar.New()
	en_uni = ut.New(en, en)
	ar_uni = ut.New(ar, ar)
	var trans ut.Translator
	trans, _ = ar_uni.GetTranslator("en")
	if "en" == "ar" {
		ar_translations.RegisterDefaultTranslations(validate, trans)
	} else {
		en_translations.RegisterDefaultTranslations(validate, trans)
	}

	err := val(data)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	return errs.Translate(trans)
}

func val[T any](data T) error {

	err := validate.Struct(data)

	return err
}
