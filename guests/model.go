package guests

// Guest struct
type Guest struct {
	Id               int    `json:"id"`
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	PhoneNumber      string `json:"phone_number" validate:"required"`
	IdentificationId string `json:"identification_id" validate:"required"`
}
