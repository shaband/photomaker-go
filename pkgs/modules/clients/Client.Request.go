package clients

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type CreateClientRequest struct {
	NameAr string                `form:"name_ar" validate:"required"`
	NameEn string                `form:"name_en" validate:"required"`
	Image  *multipart.FileHeader `form:"cover" validate:"required,image"`
}

func (ClientRequest *CreateClientRequest) ToEntity(c *gin.Context) *Client {

	return &Client{
		NameAr: ClientRequest.NameAr,
		NameEn: ClientRequest.NameEn,
		Image:  SaveImage(c, filePath, ClientRequest.Image),
	}

}

type UpdateClientRequest struct {
	NameAr string                `form:"name_ar" validate:"required"`
	NameEn string                `form:"name_en" validate:"required"`
	Image  *multipart.FileHeader `form:"image"`
}

func (ClientRequest *UpdateClientRequest) ToEntity(c *gin.Context) *Client {

	client := &Client{
		NameAr: ClientRequest.NameAr,
		NameEn: ClientRequest.NameEn,
		// Image:  SaveImage(c, filePath, UpdateClientRequest.Image),

	}
	if ClientRequest.Image != nil {
		client.Image = SaveImage(c, filePath, ClientRequest.Image)
	}

	return client
}
