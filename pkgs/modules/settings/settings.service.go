package settings

import (
	"reflect"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func (service *Service) GetAll() []Setting {
	Settings := []Setting{}
	service.db.Find(&Settings)
	return Settings
}
func (service *Service) GetAllValuesPluckedBy(key string) map[string]string {
	Settings := []Setting{}
	service.db.Find(&Settings)
	pluckedSettings := make(map[string]string)
	for _, setting := range Settings {
		pluckedSettings[getAttr(setting, key).String()] = setting.Value
	}
	return pluckedSettings
}

func (service *Service) FindValue(slug string) *string {
	Setting := Setting{}
	service.db.Select("value").Where("slug = ?", slug).Find(&Setting)
	return &(Setting.Value)
}
func NewService(db *gorm.DB) *Service {

	return &Service{
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
