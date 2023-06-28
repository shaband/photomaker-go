package contacts

import (
	"mime/multipart"
)

type ContactForm struct {
	Name             string                `form:"name" validate:"required"`
	Subject          string                `form:"subject" validate:"required"`
	Phone            string                `form:"phone" validate:"required"`
	Email            string                `form:"email" validate:"required,email"`
	ServiceTypes     []int                 `form:"service_types[]" validate:"required,dive,required" `
	ServiceTypeItems map[int]string        `form:"service_type_items" validate:"required"`
	Attachment       *multipart.FileHeader `form:"-" validate:"omitempty,file"`
}
