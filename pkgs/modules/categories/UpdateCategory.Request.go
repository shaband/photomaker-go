package categories

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type UpdateCategoryRequest struct {
	NameAr string                  `form:"name_ar" validate:"required"`
	NameEn string                  `form:"name_en" validate:"required"`
	Cover  *multipart.FileHeader   `form:"cover"`
	Images []*multipart.FileHeader `form:"images[]"`
}

func (UpdateCategoryRequest *UpdateCategoryRequest) ToEntity(c *gin.Context) *Category {

	category := &Category{
		NameAr: UpdateCategoryRequest.NameAr,
		NameEn: UpdateCategoryRequest.NameEn,
		Images: handleCategoryImages(c, UpdateCategoryRequest.Images),
	}
	if UpdateCategoryRequest.Cover != nil {
		category.Cover = SaveImage(c, filePath, UpdateCategoryRequest.Cover)
	}
	return category
}
