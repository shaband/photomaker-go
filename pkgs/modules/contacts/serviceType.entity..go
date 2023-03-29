package contacts

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type ServiceType struct {
	gorm.Model
	Order  uint
	NameAr string
	NameEn string
}

func (s *ServiceType) Fake() interface{} {

	s.Order = gofakeit.UintRange(1, 1000)
	s.NameAr = gofakeit.Word()
	s.NameEn = gofakeit.Word()

	return s
}
