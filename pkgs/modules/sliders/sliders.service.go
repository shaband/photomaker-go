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
	GetAll() *[]Slider
	Find(ID int) *Slider
	DeleteById(ID int) *gorm.DB
	Store(c *gin.Context, SliderRequest *CreateSliderRequest) *Slider
}
type Service struct {
	db DB
}
const filePath = "assets/uploads/sliders"

func NewService(db *gorm.DB) Service {
	return Service{
		db: db,
	}
}

func (s Service) GetSliders(conds ...interface{}) []*Slider {

	Sliders := []*Slider{}

	s.db.Find(&Sliders, conds...)

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
