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

func (s *ServiceTypeItem) Fake() interface{} {
	var types []string = []string{
		"text", "checkbox", "number",
	}

	s.NameAr = gofakeit.Word()
	s.NameEn = gofakeit.Word()
	s.InputType = gofakeit.RandomString(types)

	return s
}
