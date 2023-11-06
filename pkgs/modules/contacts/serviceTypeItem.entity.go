package contacts

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type ServiceTypeItem struct {
	gorm.Model
	NameAr        string
	NameEn        string
	InputType     string
	ServiceTypeId int
	ServiceType   ServiceType
}

func (s *ServiceTypeItem) Fake() interface{} {
	var types []string = []string{
		"text", "checkbox", "number",
	}

	ServiceType := ServiceType{}
	ServiceType.Fake()
	s.NameAr = gofakeit.Word()
	s.NameEn = gofakeit.Word()
	s.InputType = gofakeit.RandomString(types)
	s.ServiceType = ServiceType

	return s
}
