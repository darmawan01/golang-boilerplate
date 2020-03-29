package order_guests

// OrderGuest struct
type OrderGuest struct {
	Id          int    `json:"id"`
	OrderId     int    `json:"order_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Email       int    `json:"email" validate:"required"`
	PhoneNumber int    `json:"phone_number" validate:"required"`
}
