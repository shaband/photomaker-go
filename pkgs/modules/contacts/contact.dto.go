package contacts

import "mime/multipart"

type ContactForm struct {
	Name             string                `form:"name" binding:"required"`
	Subject          string                `form:"subject" binding:"required"`
	Phone            string                `form:"phone" binding:"required"`
	Email            string                `form:"email" binding:"required,email"`
	ServiceTypes     []int                 `form:"service_types[]" binding:"required"`
	ServiceTypeItems map[int][]string      `form:"service_type_items"`
	Attachment       *multipart.FileHeader `form:"attachment,omitempty"`
}
