package hotels

// Hotel struct
type Hotel struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Latitute  string `json:"latitute" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
}

type roomRates struct {
	Price float64 `json:"price"`
	Date  string  `json:"date"`
}

type hotelRoom struct {
	RoomName   string      `json:"room"`
	RoomPrices []roomRates `json:"prices"`
}

type hotelRates struct {
	HotelName string      `json:"hotel"`
	Rooms     []hotelRoom `json:"rooms"`
}
