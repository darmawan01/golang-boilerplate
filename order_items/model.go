package order_items

// OrderItem struct
type OrderItem struct {
	Id       int     `json:"id"`
	OrderId  int     `json:"order_id"`
	RoomId   int     `json:"room_id" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}
