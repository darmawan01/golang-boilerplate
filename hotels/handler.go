package hotels

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jackc/pgx"
)

func (api *HotelsApi) allHandler() (hotels []Hotel, err error) {
	var rows *pgx.Rows
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

func (api *HotelsApi) bodyToStruct(body io.ReadCloser) (model Hotel, err error) {
	if err = json.NewDecoder(body).Decode(&model); err != nil {
		return
	}
	return
}
