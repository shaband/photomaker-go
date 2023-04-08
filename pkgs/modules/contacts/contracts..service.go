package contacts

import "gorm.io/gorm"

type contractService struct {
	db *gorm.DB
}

func (service *contractService) GetAll() *[]*ServiceType {
	ServiceTypes := []*ServiceType{}
	service.db.Preload("Items").Find(&ServiceTypes)
	return &ServiceTypes
}

func NewContractService(db *gorm.DB) *contractService {

	return &contractService{
		db: db,
	}
}
