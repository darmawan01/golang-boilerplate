package hotels

// Hotel struct
type Hotel struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Latitute  string `json:"latitute" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
}
