package categories

type CreateCategoryRequest struct {
	NameAr string          `form:"name_ar" validate:"required,alphanum"`
	NameEn string          `form:"name_en" validate:"required,alphanum"`
	Cover  string          `form:"cover" validate:"required"`
	Images []CategoryImage `form:"images[]" validate:"required"`
}

func (c *CreateCategoryRequest) ToEntity() *Category {

	return &Category{
		NameAr: c.NameAr,
		NameEn: c.NameEn,
		Cover:  c.Cover,
		Images: c.Images,
	}

}
