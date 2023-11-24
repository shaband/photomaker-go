package settings

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Select(value interface{}) *gorm.DB
}

type ServiceInterface interface {
	All(conds ...interface{}) []*Setting
	Find(ID int) *Setting
	Update(ctx *gin.Context, ID uint, Request *SettingRequest) *gorm.DB
	GetAllValuesPluckedBy(key string) map[string]string
	FindValue(slug string) *string
}

func (service Service) All(conds ...interface{}) []*Setting {

	settings := []*Setting{}

	service.db.Find(&settings)
	return settings
}
func (service Service) GetAll() []Setting {
	Settings := []Setting{}
	service.db.Find(&Settings)
	return Settings
}
func (service Service) GetAllValuesPluckedBy(key string) map[string]string {
	Settings := []Setting{}
	service.db.Find(&Settings)
	pluckedSettings := make(map[string]string)
	for _, setting := range Settings {
		pluckedSettings[getAttr(setting, key).String()] = setting.Value
	}
	return pluckedSettings
}

func (service Service) Find(ID int) *Setting {
	setting := Setting{}
	service.db.Find(&setting, ID)
	return &setting
}
func (service Service) Update(ctx *gin.Context, ID uint, setting *SettingRequest) *gorm.DB {
	entity := setting.ToEntity()
	entity.ID = ID
	return service.db.Model(entity).Updates(entity)

}

func (service Service) DeleteById(ID int) *Service {
	service.db.Delete(&Service{}, ID)
	return nil
}

func (service Service) FindValue(slug string) *string {
	Setting := Setting{}
	service.db.Select("value").Where("slug = ?", slug).Find(&Setting)
	return &(Setting.Value)
}
func NewService(db *gorm.DB) Service {

	return Service{
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
