package contacts

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type ServiceTypeItem struct {
	gorm.Model
	NameAr    string
	NameEn    string
	InputType string
}

func (s ServiceTypeItem) Fake() *ServiceTypeItem {
	var types []string = []string{
		"text", "checkbox", "number",
	}

	s = ServiceTypeItem{
		NameAr:    gofakeit.Word(),
		NameEn:    gofakeit.Word(),
		InputType: gofakeit.RandomString(types),
	}
	return &s
}
