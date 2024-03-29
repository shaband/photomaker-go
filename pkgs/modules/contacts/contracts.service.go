package contacts

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"golang.org/x/exp/maps"
	gomail "gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

const filePath = "assets/uploads/contacts"

func (service *Service) All() *[]*Contact {
	contacts := []*Contact{}
	service.db.Find(&contacts)
	return &contacts
}
func (service *Service) GetAll() *[]*ServiceType {
	ServiceTypes := []*ServiceType{}
	service.db.Preload("Items").Find(&ServiceTypes)
	return &ServiceTypes
}

func (service *Service) BindForm(context *gin.Context) ContactForm {

	var form ContactForm

	if err := context.ShouldBind(&form); err != nil {
		// handle error
		panic(err)

	}

	services := form.ServiceTypes

	for _, service := range services {
		if form.ServiceTypeItems == nil {
			form.ServiceTypeItems = make(map[string]string)
		}
		
		form.ServiceTypeItems[fmt.Sprint(service)] = context.PostForm(fmt.Sprintf("service_type_items[%v]", service))
	}

	form.Attachment, _ = context.FormFile("attachment")
	return form
}

func NewService(db *gorm.DB) *Service {

	return &Service{
		db: db,
	}
}

func (service *Service) SaveContactData(c *gin.Context, data ContactForm) *Contact {
	var dest string
	if data.Attachment != nil {
		dest = helpers.SaveFile(c, filePath, data.Attachment)
	}
	contact := Contact{
		Name:         data.Name,
		Phone:        data.Phone,
		Email:        data.Email,
		File:         dest,
		Category:     data.Subject,
		ServiceTypes: data.ServiceTypeItems,
	}
	result := service.db.Create(&contact)
	fmt.Println(result)
	return &contact
}
func (service *Service) Update(ctx *gin.Context, ID uint, Request *ContactRequest) *gorm.DB {
	entity := Request.ToEntity()
	entity.ID = ID

	trx :=service.db.Model(entity).Updates(entity)
	trx.Find(&entity,ID)

	msg := gomail.NewMessage()
    msg.SetHeader("From", "photomaker@photomaker.com")
    msg.SetHeader("To", entity.Email)
    msg.SetHeader("Subject", "replying on your message to phomaker ")
    msg.SetBody("text/html",fmt.Sprintf("hello %s </br> %s",entity.Name, entity.Replay))

    n := gomail.NewDialer("localhost", 1025, "", "")

    // Send the email
    if err := n.DialAndSend(msg); err != nil {
        log.Println(err)
    }
	return trx
}

func (service *Service) Find(ID int) *Contact {
	Contact := Contact{}
	service.db.Find(&Contact, ID)
	SelectedserviceTypeItems:=Contact.ServiceTypes;
	servicesIDs := maps.Keys(SelectedserviceTypeItems)
	services := []ServiceType{}
	service.db.Find(&services, servicesIDs)
m:=make(map[string]string)
	for _, serviceType := range services {
		m[serviceType.NameEn]=SelectedserviceTypeItems[fmt.Sprint(serviceType.ID)]
	}
	Contact.ServiceTypes=m
	return &Contact
}
