package sliders

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type SliderRequest struct {
	Image *multipart.FileHeader `form:"cover" validate:"required,image"`
}

func (SliderRequest *SliderRequest) ToEntity(c *gin.Context) *Slider {

	return &Slider{
		Image: SaveImage(c, filePath, SliderRequest.Image),
	}

}
