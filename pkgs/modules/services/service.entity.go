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

func (s *Service) Fake() interface{} {

	s.TitleAr = gofakeit.Word()
	s.TitleEn = gofakeit.Word()
	s.DescriptionAr = gofakeit.Paragraph(1, 3, 5, " ")
	s.DescriptionEn = gofakeit.Paragraph(1, 3, 5, " ")

	return s
}
