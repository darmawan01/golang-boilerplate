package orders

import (
	orderGuests "kodingworks/order_guests"
	orderItems "kodingworks/order_items"
	"time"
)

type Status string

const (
	Pending   Status = "pending"
	Paid      Status = "paid"
	Ready     Status = "ready"
	Checkin   Status = "checkin"
	Checkout  Status = "checkout"
	Expired   Status = "expired"
	Cancelled Status = "cancelled"
)

// Order struct
type Order struct {
	Id         int       `json:"id"`
	HotelId    int       `json:"hotel_id" validate:"required"`
	GuestId    int       `json:"guest_id" validate:"required"`
	Status     string    `json:"status" validate:"required"`
	CheckinAt  time.Time `json:"checkin_at" validate:"required"`
	CheckoutAt time.Time `json:"checkout_at" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderRequests struct {
	HotelId     int                    `json:"hotel_id" validate:"required"`
	GuestId     int                    `json:"guest_id" validate:"required"`
	Status      string                 `json:"status" validate:"required"`
	CheckinAt   time.Time              `json:"checkin_at" validate:"required"`
	CheckoutAt  time.Time              `json:"checkout_at" validate:"required"`
	OrderItems  []orderItems.OrderItem `json:"order_items"`
	OrderGuests orderGuests.OrderGuest `json:"order_guests"`
}

type report struct {
	TotalOrders int                    `json:"total_orders"`
	TotalSales  float64                `json:"total_sales"`
	Totals      map[string]interface{} `json:"totals"`
}
