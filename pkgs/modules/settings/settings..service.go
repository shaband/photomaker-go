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

func (service *SettingService) FindValue(slug string) *string {
	Setting := Setting{}
	service.db.Select("value").Where("slug = ?", slug).Find(&Setting)
	return &(Setting.Value)
}
func NewSettingService(db *gorm.DB) *SettingService {

	return &SettingService{
		db: db,
	}
}
