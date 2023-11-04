package categories

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type CreateCategoryRequest struct {
	NameAr string                  `form:"name_ar" validate:"required"`
	NameEn string                  `form:"name_en" validate:"required"`
	Cover  *multipart.FileHeader   `form:"cover" validate:"required,image"`
	Images []*multipart.FileHeader `form:"images[]" validate:"required"`
}

func (CreateCategoryRequest *CreateCategoryRequest) ToEntity(c *gin.Context) *Category {

	return &Category{
		NameAr: CreateCategoryRequest.NameAr,
		NameEn: CreateCategoryRequest.NameEn,
		Cover:  SaveImage(c, filePath, CreateCategoryRequest.Cover),
		Images: handleCategoryImages(c, CreateCategoryRequest.Images),
	}

}
