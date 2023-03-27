package categories

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	NameAr string
	NameEn string
	Cover  string
}

func (c Category) Fake() *Category {

	c = Category{
		NameAr: gofakeit.Name(),
		NameEn: gofakeit.Name(),
		Cover:  gofakeit.ImageURL(400, 400),
	}
	return &c
}
