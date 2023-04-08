package clients

import "gorm.io/gorm"

type clientService struct {
	db *gorm.DB
}

func NewClientSerive(db *gorm.DB) clientService {

	return clientService{
		db: db,
	}
}

func (service clientService) All() []*Client {

	Clients := []*Client{}

	service.db.Find(&Clients)
	return Clients
}
