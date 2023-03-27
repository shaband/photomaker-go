package services

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	TitleAr       string
	DescriptionAr string
	TitleEn       string
	DescriptionEn string
}

func (s Service) Fake() *Service {
	s = Service{
		TitleAr:       gofakeit.Word(),
		TitleEn:       gofakeit.Word(),
		DescriptionAr: gofakeit.Paragraph(1, 3, 5, " "),
		DescriptionEn: gofakeit.Paragraph(1, 3, 5, " "),
	}
	return &s
}
