package clients

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	NameAr string
	NameEn string
	Image  string
}
