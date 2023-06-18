package contacts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type contractService struct {
	db *gorm.DB
}

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

	return form
}

func NewContractService(db *gorm.DB) *contractService {

	return &contractService{
		db: db,
	}
}
