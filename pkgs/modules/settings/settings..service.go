package settings

import "gorm.io/gorm"

type SettingService struct {
	db *gorm.DB
}

func (service *SettingService) GetAll() []Setting {
	Settings := []Setting{}
	service.db.Find(&Settings)
	return Settings
}

func (service *SettingService) Find(slug string) *Setting {
	Settings := Setting{
		Slug: slug,
	}
	service.db.Find(&Settings)
	return &Settings
}

func NewSettingService(db *gorm.DB) *SettingService {

	return &SettingService{
		db: db,
	}
}
