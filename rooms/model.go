package rooms

type Room struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	HotelId  int     `json:"hotel_id" validate:"required"`
}
