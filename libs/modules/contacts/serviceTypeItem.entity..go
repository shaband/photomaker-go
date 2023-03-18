package contacts

import "gorm.io/gorm"

type ServiceTypeItem struct{

	gorm.Model
	NameAr string
	NameEn string
	InputType string
}