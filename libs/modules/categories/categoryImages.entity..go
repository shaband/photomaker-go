package categories

import "gorm.io/gorm"

type CategoryImage struct {
	gorm.Model
	Image string
	Order uint
}
