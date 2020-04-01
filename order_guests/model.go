package order_guests

// OrderGuest struct
type OrderGuest struct {
	Id          int    `json:"id"`
	OrderId     int    `json:"order_id"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
