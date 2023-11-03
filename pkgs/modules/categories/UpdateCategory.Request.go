package categories

import "mime/multipart"

type UpdateCategoryRequest struct {
	NameAr string                  `form:"name_ar" validate:"required,alphanum"`
	NameEn string                  `form:"name_en" validate:"required,alphanum"`
	Cover  *multipart.FileHeader   `form:"cover"`
	Images []*multipart.FileHeader `form:"images[]"`
}

func (c *UpdateCategoryRequest) ToEntity() *Category {

	return &Category{
		NameAr: c.NameAr,
		NameEn: c.NameEn,
		//Cover:  c.Cover,
		//Images: c.Images,
	}

}
