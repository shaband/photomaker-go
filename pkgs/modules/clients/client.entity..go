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

func (c Client) Fake() *Client {
	c = Client{
		NameAr: gofakeit.Name(),
		NameEn: gofakeit.Name(),
		Image:  gofakeit.ImageURL(400, 400),
	}
	return &c

}
