package sliders

import (
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}

type ServiceInterface interface {
	All(conds ...interface{}) []*Slider
	Find(ID int) *Slider
	DeleteById(ID int) *gorm.DB
	Store(c *gin.Context, SliderRequest *SliderRequest) *Slider
	Update(ctx *gin.Context, ID uint, SliderRequest *SliderRequest) *gorm.DB
}
type Service struct {
	db DB
}

const filePath = "assets/uploads/sliders"

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (service *Service) GetSliders(conds ...interface{}) []*Slider {

	Sliders := []*Slider{}

	service.db.Find(&Sliders, conds...)

	return Sliders
}

func SaveImage(c *gin.Context, dest string, image *multipart.FileHeader) string {
	path := dest + "/" + fmt.Sprint(time.Now().Unix()) + "/" + image.Filename

	err := c.SaveUploadedFile(image, path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return "/" + path
}

func (service *Service) All(conds ...interface{}) []*Slider {

	Sliders := []*Slider{}

	service.db.Find(&Sliders)
	return Sliders
}

func (service *Service) Find(ID int) *Slider {
	Slider := Slider{}
	service.db.Preload("Images").Find(&Slider, ID)

	return &Slider
}
func (service *Service) Update(ctx *gin.Context, ID uint, SliderRequest *SliderRequest) *gorm.DB {
	category := SliderRequest.ToEntity(ctx)
	category.ID = ID
	return service.db.Save(category)
}

func (service *Service) DeleteById(ID int) *gorm.DB {
	// Slider:=service.Find(ID)
	return service.db.Delete(&Slider{}, ID)

}

func (service *Service) Store(c *gin.Context, SliderRequest *SliderRequest) *Slider {
	category := SliderRequest.ToEntity(c)
	service.db.Create(category)
	return category
}
