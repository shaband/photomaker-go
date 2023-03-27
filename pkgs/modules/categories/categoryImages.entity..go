package categories

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type CategoryImage struct {
	gorm.Model
	Image string `fake:"{imageURL:400x400}"`
	Order uint   `fake:"{Number:1,100}"`
}

func (c CategoryImage) Fake() *CategoryImage {
	c = CategoryImage{
		Order: uint(gofakeit.Number(1, 1000)),
		Image: gofakeit.ImageURL(400, 400),
	}
	return &c
}
