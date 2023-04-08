package services

import "gorm.io/gorm"

type servicesService struct {
	db *gorm.DB
}

func NewServicesSerive(db *gorm.DB) servicesService {

	return servicesService{
		db: db,
	}
}

func (service servicesService) All() []*Service {
	services := []*Service{}
	service.db.Find(&services)
	return services
}
