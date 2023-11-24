package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type serviceS struct {
	db *gorm.DB
}

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}

type ServiceInterface interface {
	All(conds ...interface{}) []*Service
	Find(ID int) *Service
	DeleteById(ID int) *gorm.DB
	Store(c *gin.Context, ServiceRequest *ServiceRequest) *Service
	Update(ctx *gin.Context, ID uint, ServiceRequest *ServiceRequest) *gorm.DB
}

func NewService(db *gorm.DB) *serviceS {

	return &serviceS{
		db: db,
	}
}

func (service serviceS) All(conds ...interface{}) []*Service {
	services := []*Service{}
	service.db.Find(&services)
	return services
}

func (service *serviceS) GetAll() *[]Service {
	Services := []Service{}
	service.db.Find(&Services)
	return &Services
}
func (service *serviceS) FindServiceByEmail(Email string) *Service {
	Service := Service{}
	service.db.Where("email = ?", Email).Find(&Service)
	return &Service
}

func (service *serviceS) Find(ID int) *Service {
	Service := Service{}
	service.db.Find(&Service, ID)
	return &Service
}
func (service serviceS) Update(ctx *gin.Context, ID uint, ServiceRequest *ServiceRequest) *gorm.DB {

	entity := ServiceRequest.ToEntity()
	entity.ID = ID
	return service.db.Save(entity)
}

func (service *serviceS) DeleteById(ID int) *gorm.DB {
	// user:=service.Find(ID)
	return service.db.Delete(&Service{}, ID)

}

func (service *serviceS) Store(ctx *gin.Context, request *ServiceRequest) *Service {

	entity := request.ToEntity()
	service.db.Create(entity)

	return entity
}
