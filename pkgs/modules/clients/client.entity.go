package clients

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model

	NameAr string
	NameEn string
	Image  string
}

func (c *Client) Fake() interface{} {
	c.NameAr = gofakeit.Name()
	c.NameEn = gofakeit.Name()
	c.Image = gofakeit.ImageURL(400, 400)

	return c

}
