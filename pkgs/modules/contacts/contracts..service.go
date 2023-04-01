package contacts

import "gorm.io/gorm"

type ContractService struct {
	db *gorm.DB
}

func (service *ContractService) GetAll() *[]*ServiceType {
	ServiceTypes := []*ServiceType{}
	service.db.Preload("Items").Find(&ServiceTypes)
	return &ServiceTypes
}

func NewContractService(db *gorm.DB) *ContractService {

	return &ContractService{
		db: db,
	}
}
