package room_availabalities

import "time"

// RoomAvailabilities struct
type RoomAvailabilities struct {
	Id       int       `json:"id"`
	RoomId   int       `json:"room_id" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	Quantity int       `json:"quantity" validate:"required"`
}
