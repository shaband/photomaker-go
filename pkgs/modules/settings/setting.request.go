package settings

type SettingRequest struct {
	Value string `form:"value" validate:"required"`
}

func (Request *SettingRequest) ToEntity() *Setting {

	return &Setting{
		Value: Request.Value,
	}
}
