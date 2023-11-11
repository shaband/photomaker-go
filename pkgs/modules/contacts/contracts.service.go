package contacts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"gorm.io/gorm"
)

type contractService struct {
	db *gorm.DB
}

const filePath = "assets/uploads/contacts"

func (service *contractService) GetAll() *[]*ServiceType {
	ServiceTypes := []*ServiceType{}
	service.db.Preload("Items").Find(&ServiceTypes)
	return &ServiceTypes
}

func (service *contractService) BindForm(context *gin.Context) ContactForm {

	var form ContactForm

	if err := context.ShouldBind(&form); err != nil {
		// handle error
		panic(err)

	}

	services := form.ServiceTypes

	for _, service := range services {
		if form.ServiceTypeItems == nil {
			form.ServiceTypeItems = make(map[int]string)
		}
		form.ServiceTypeItems[service] = context.PostForm(fmt.Sprintf("service_type_items[%v]", service))
	}

	form.Attachment, _ = context.FormFile("attachment")
	return form
}

func NewService(db *gorm.DB) *contractService {

	return &contractService{
		db: db,
	}
}

func (service *contractService) SaveContactData(c *gin.Context, data ContactForm) *Contact {
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
