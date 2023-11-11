package sliders

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type CreateSliderRequest struct {
	Image  *multipart.FileHeader `form:"cover" validate:"required,image"`
}

func (SliderRequest *CreateSliderRequest) ToEntity(c *gin.Context) *Slider {

	return &Slider{
		Image:  SaveImage(c, filePath, SliderRequest.Image),
	}

}

// type UpdateSliderRequest struct {
// 	Image  *multipart.FileHeader `form:"image"`
// }

// func (SliderRequest *UpdateSliderRequest) ToEntity(c *gin.Context) *Slider {

// 	client := &Slider{
// 		// Image:  SaveImage(c, filePath, UpdateSliderRequest.Image),

// 	}
// 	if SliderRequest.Image != nil {
// 		client.Image = SaveImage(c, filePath, SliderRequest.Image)
// 	}

// 	return client
// }
