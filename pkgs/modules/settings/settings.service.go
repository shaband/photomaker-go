package settings

import (
	"fmt"
	"log"
	"mime/multipart"
	"reflect"
	"time"

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
	Update(ctx *gin.Context, ID uint) (error, *gorm.DB)
	GetAllValuesPluckedBy(key string) map[string]string
	FindValue(slug string) *string
}

func (service Service) All(conds ...interface{}) []*Setting {

	settings := []*Setting{}

	service.db.Find(&settings)
	return settings
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
func (service Service) Update(ctx *gin.Context, ID uint) (error, *gorm.DB) {
	entity := Setting{}
	trx := service.db.Find(&entity, ID)
	if trx.Error != nil {
		return trx.Error, nil
	}

	switch entity.Type {
	case "file":
		file, err := ctx.FormFile("value")
		if err != nil {
			return err, nil
		}
		entity.Value = saveFile(ctx, filePath, file)
		break
	// case "text":
	// 	entity.Value = ctx.PostForm("value")
	// 	break
	default:
		entity.Value = ctx.PostForm("value")
		break
	}

	if entity.Value == "" {
		return fmt.Errorf("value is required"), nil
	}

	trx = service.db.Save(&entity)
	if trx.Error != nil {
		return trx.Error, nil
	}
	return nil, trx
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

const filePath = "assets/uploads/clients"

func saveFile(c *gin.Context, dest string, file *multipart.FileHeader) string {
	path := dest + "/" + fmt.Sprint(time.Now().Unix()) + "/" + file.Filename

	err := c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return "/" + path
}
