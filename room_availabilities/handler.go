package room_availabalities

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jackc/pgx"
)

func (api *RoomAvailabilitiesApi) allHandler() (roomAvailabilities []RoomAvailabilities, err error) {
	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM room_availabilities`); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var roomAvailability RoomAvailabilities
		if err = rows.Scan(
			&roomAvailability.Id,
			&roomAvailability.RoomId,
			&roomAvailability.Date,
			&roomAvailability.Quantity,
		); err != nil {
			log.Println("allHandler(): ", err.Error())
			if err == pgx.ErrNoRows {
				roomAvailabilities, err = []RoomAvailabilities{}, nil
				return
			}
			return
		}
		roomAvailabilities = append(roomAvailabilities, roomAvailability)
	}
	return
}

func (api *RoomAvailabilitiesApi) addHandler(roomAvailability RoomAvailabilities) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO room_availabilities (date, quantity, room_id) VALUES($1, $2, $3) RETURNING id`,
		roomAvailability.Date, roomAvailability.Quantity, roomAvailability.RoomId).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}

func (api *RoomAvailabilitiesApi) detailHandler(id int) (roomAvailability RoomAvailabilities, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM room_availabilities WHERE id=$1`, id).Scan(
		&roomAvailability.Id,
		&roomAvailability.RoomId,
		&roomAvailability.Date,
		&roomAvailability.Quantity,
	); err != nil {
		log.Println("detailHandler(): ", err.Error())
		if err == pgx.ErrNoRows {
			roomAvailability, err = RoomAvailabilities{}, nil
			return
		}
		return
	}
	return
}

func (api *RoomAvailabilitiesApi) updateHandler(roomAvailability RoomAvailabilities) (err error) {
	var result pgx.CommandTag
	result, err = api.Db.Exec(`
		UPDATE room_availabilities SET date=$1, quantity=$2, room_id=$3 WHERE id=$4`,
		roomAvailability.Date,
		roomAvailability.Quantity,
		roomAvailability.RoomId,
		roomAvailability.Id,
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

func (api *RoomAvailabilitiesApi) deleteHandler(id int) (err error) {
	var result pgx.CommandTag
	if result, err = api.Db.Exec(`DELETE FROM room_availabilities WHERE id=$1`, id); err != nil {
		log.Println("deleteHandler(): ", err.Error())
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-delete")
		return
	}
	return
}

func (api *RoomAvailabilitiesApi) bodyToStruct(body io.ReadCloser) (model RoomAvailabilities, err error) {
	if err = json.NewDecoder(body).Decode(&model); err != nil {
		return
	}
	return
}
