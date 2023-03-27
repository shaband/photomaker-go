package contacts

import "github.com/brianvoe/gofakeit/v6"

type ServiceType struct {
	Order  uint
	NameAr string
	NameEn string
}

func (s ServiceType) Fake() *ServiceType {
	s = ServiceType{
		Order:  uint(gofakeit.Number(1, 1000)),
		NameAr: gofakeit.Word(),
		NameEn: gofakeit.Word(),
	}
	return &s
}
