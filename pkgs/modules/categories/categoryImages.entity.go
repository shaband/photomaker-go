package categories

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type CategoryImage struct {
	gorm.Model
	Image    string
	Order    uint
	CategoryId int
	Category Category  
}

func (c *CategoryImage) Fake() interface{} {

	category := Category{}
	category.Fake()
	c.Order = gofakeit.UintRange(1, 1000)
	c.Image = gofakeit.ImageURL(400, 400)
	c.Category = category
	return c
}
