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
	Images []CategoryImage
}

func (c *Category) Fake() interface{} {

	c.NameAr = gofakeit.Name()
	c.NameEn = gofakeit.Name()
	c.Cover = gofakeit.ImageURL(400, 400)

	return c
}
