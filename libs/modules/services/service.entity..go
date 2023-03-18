package services

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	TitleAr       string
	DescriptionAr string
	TitleEr       string
	DescriptionEr string
}
