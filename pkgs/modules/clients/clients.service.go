package clients

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
	All(conds ...interface{}) []*Client
	GetAll() *[]Client
	Find(ID int) *Client
	Update(ctx *gin.Context, ID uint, ClientRequest *UpdateClientRequest) *gorm.DB
	DeleteById(ID int) *gorm.DB
	Store(c *gin.Context, ClientRequest *CreateClientRequest) *Client
}
type Service struct {
	db DB
}

const filePath = "assets/uploads/clients"

func NewService(db *gorm.DB) Service {

	return Service{
		db: db,
	}
}

func (service Service) All() []*Client {

	Clients := []*Client{}

	service.db.Find(&Clients)
	return Clients
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

func (service *Service) Find(ID int) *Client {
	Client := Client{}
	service.db.Preload("Images").Find(&Client, ID)

	return &Client
}
func (service Service) Update(ctx *gin.Context, ID uint, ClientRequest *UpdateClientRequest) *gorm.DB {
	category := ClientRequest.ToEntity(ctx)
	category.ID = ID
	return service.db.Save(category)
}

func (service *Service) DeleteById(ID int) *gorm.DB {
	// Client:=service.Find(ID)
	return service.db.Delete(&Client{}, ID)

}

func (service *Service) Store(c *gin.Context, ClientRequest *CreateClientRequest) *Client {
	category := ClientRequest.ToEntity(c)
	service.db.Create(category)
	return category
}

