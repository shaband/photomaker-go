package settings

import (
	"reflect"

	"gorm.io/gorm"
)

type SettingService struct {
	db *gorm.DB
}

func (service *SettingService) GetAll() []Setting {
	Settings := []Setting{}
	service.db.Find(&Settings)
	return Settings
}
func (service *SettingService) GetAllValuesPluckedBy(key string) map[string]string {
	Settings := []Setting{}
	service.db.Find(&Settings)
	pluckedSettings := make(map[string]string)
	for _, setting := range Settings {
		pluckedSettings[getAttr(setting, key).String()] = setting.Value
	}
	return pluckedSettings
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

func getAttr(obj Setting, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable

	curField := pointToStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}

	return curField
}
