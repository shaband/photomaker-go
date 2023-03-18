package sliders

import "gorm.io/gorm"

type Slider struct{
 gorm.Model
 Order uint
 Image string

}