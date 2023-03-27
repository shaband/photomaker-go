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

func (c *CategoryImage) Fake() interface{} {
	
		c.Order= gofakeit.UintRange(1,1000);
		c.Image= gofakeit.ImageURL(400, 400);
	
	return c
}
