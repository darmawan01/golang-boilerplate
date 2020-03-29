package room_rates

import "time"

type RoomRate struct {
	Id     int       `json:"id"`
	RoomId int       `json:"room_id" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
	Price  float64   `json:"price" validate:"required"`
}
