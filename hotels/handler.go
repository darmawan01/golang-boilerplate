package hotels

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jackc/pgx"
)

var rows *pgx.Rows

func (api *HotelsApi) allHandler() (hotels []Hotel, err error) {
	if rows, err = api.Db.Query(`SELECT * FROM hotels`); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var hotel Hotel

		if err = rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Latitute, &hotel.Longitude); err != nil {
			log.Println("allHandler(): ", err.Error())
			if err == pgx.ErrNoRows {
				hotels, err = []Hotel{}, nil
				return
			}
			return
		}
		hotels = append(hotels, hotel)
	}
	return
}

func (api *HotelsApi) addHandler(hotel Hotel) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO hotels (name, address, latitude, longitude) VALUES ($1, $2, $3, $4) RETURNING id`,
		hotel.Name, hotel.Address, hotel.Latitute, hotel.Longitude).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}

func (api *HotelsApi) detailHandler(id int) (hotel Hotel, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM hotels WHERE id=$1`, id).Scan(
		&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Latitute, &hotel.Longitude); err != nil {
		log.Println("detailHandler(): ", err.Error())
		if err == pgx.ErrNoRows {
			hotel, err = Hotel{}, nil
			return
		}
		return
	}
	return
}

func (api *HotelsApi) updateHandler(hotel Hotel) (err error) {
	var result pgx.CommandTag
	result, err = api.Db.Exec(`
		UPDATE hotels SET name=$1, address=$2, latitude=$3, longitude=$4 WHERE id=$5`,
		hotel.Name,
		hotel.Address,
		hotel.Latitute,
		hotel.Longitude,
		hotel.Id,
	)
	if err != nil {
		log.Println("updateHandler(): ", err.Error())
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-update")
		return
	}
	return
}

func (api *HotelsApi) deleteHandler(id int) (err error) {
	var result pgx.CommandTag
	if result, err = api.Db.Exec(`DELETE FROM hotels WHERE id=$1`, id); err != nil {
		log.Println("deleteHandler(): ", err.Error())
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-delete")
		return
	}
	return
}

func (api *HotelsApi) roomRatesByHotelHandler(id int) (hotelResponse hotelRates, err error) {
	// Get the hotel detail
	hotel, err := api.detailHandler(id)
	if err != nil {
		return
	}

	getRatesPrices := func(roomId int) (ratePrices []roomRates, err error) {
		if rows, err = api.Db.Query(`
			SELECT room.price AS room_price, rate.Price, rate.date 
			FROM room_rates rate
			JOIN (SELECT id, price FROM rooms) AS room 
			ON room.id=rate.room_id WHERE room.id=$1`, roomId); err != nil {
			log.Println("getRatesPrices(): ", err.Error())
			return
		}
		defer rows.Close()

		for rows.Next() {
			var rate roomRates
			roomPrice := sql.NullFloat64{}
			var priceDate time.Time
			if err = rows.Scan(
				&roomPrice,
				&rate.Price,
				&priceDate,
			); err != nil && err != pgx.ErrNoRows {
				log.Println("getRatesPrices(): ", err.Error())
				return
			}

			rate.Date = priceDate.Local().Format("01/01/2006")
			// Replace room rates price if is not set lower then 0
			if rate.Price <= 0 {
				rate.Price = roomPrice.Float64
			}
			ratePrices = append(ratePrices, rate)
		}

		return
	}

	if rows, err = api.Db.Query(`SELECT id, name FROM rooms WHERE hotel_id=$1`,
		hotel.Id); err != nil {
		log.Println("roomRatesByHotelHandler(): ", err.Error())
		return
	}
	defer rows.Close()

	rooms := make([]hotelRoom, 0)
	for rows.Next() {
		var room hotelRoom
		var roomID int

		if err = rows.Scan(&roomID, &room.RoomName); err != nil && err != pgx.ErrNoRows {
			log.Println("roomRatesByHotelHandler(): ", err.Error())
			return
		}
		var rates []roomRates
		if rates, err = getRatesPrices(roomID); err != nil {
			return
		}
		room.RoomPrices = rates
		rooms = append(rooms, room)
	}
	hotelResponse.HotelName = hotel.Name
	hotelResponse.Rooms = rooms

	return
}

func (api *HotelsApi) bodyToStruct(body io.ReadCloser) (model Hotel, err error) {
	if err = json.NewDecoder(body).Decode(&model); err != nil {
		return
	}
	return
}
