package services

type ServiceRequest struct{


	TitleAr       string `form:"title_ar" validate:"required"`
	DescriptionAr string `form:"description_ar" validate:"required"`
	TitleEn       string `form:"title_en" validate:"required"`
	DescriptionEn string `form:"description_en" validate:"required"`

}

func (serviceRequest *ServiceRequest) ToEntity() *Service {

	return &Service{
		TitleAr:       serviceRequest.TitleAr,
		DescriptionAr: serviceRequest.DescriptionAr,
		TitleEn:       serviceRequest.TitleEn,
		DescriptionEn: serviceRequest.DescriptionEn,
	}
}