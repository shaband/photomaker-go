package contacts

type ContactRequest struct {
	Name             string   `form:"name"`
	Subject          string   `form:"subject"`
	Phone            string   `form:"phone"`
	Email            string   `form:"email"`
	ServiceType      []string `form:"Service_types[]"`
	ServiceTypeItems []string `form:"Service_type_items"`
}
