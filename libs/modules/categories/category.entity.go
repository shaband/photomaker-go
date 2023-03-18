package categories

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	
	NameAr string
	NameEn string
	
	Cover string
	
  }