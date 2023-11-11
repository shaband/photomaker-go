package services

import "gorm.io/gorm"

type serviceS struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *serviceS {

	return &serviceS{
		db: db,
	}
}

func (service serviceS) All() []*Service {
	services := []*Service{}
	service.db.Find(&services)
	return services
}

