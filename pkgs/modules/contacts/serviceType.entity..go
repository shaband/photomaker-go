package contacts

import "github.com/brianvoe/gofakeit/v6"

type ServiceType struct {
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
