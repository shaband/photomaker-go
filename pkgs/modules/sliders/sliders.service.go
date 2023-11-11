package sliders

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{
		db: db,
	}
}

func (s Service) GetSliders(conds ...interface{}) []*Slider {

	Sliders := []*Slider{}

	s.db.Find(&Sliders, conds...)

	return Sliders
}
