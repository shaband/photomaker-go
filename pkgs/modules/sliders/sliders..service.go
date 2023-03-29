package sliders

import "gorm.io/gorm"

type SliderService struct {
	db *gorm.DB
}

func NewSliderService(db *gorm.DB) SliderService {
	return SliderService{
		db: db,
	}
}

func (s SliderService) GetSliders(conds ...interface{}) []*Slider {

	Sliders := []*Slider{}

	s.db.Find(&Sliders, conds...)

	return Sliders
}
