package sliders

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

type Slider struct {
	gorm.Model
	Order uint
	Image string
}
func (s *Slider) Fake() interface{} {
    s.Order = gofakeit.UintRange(1, 1000)
    s.Image = gofakeit.ImageURL(400, 400)
    return s
}