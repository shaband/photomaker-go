package contacts

type ContactRequest struct {
	Replay string `form:"replay" validate:"required,"`

}

func (u *ContactRequest) ToEntity() *Contact {

	return &Contact{
		Replay: u.Replay,
	}

}
