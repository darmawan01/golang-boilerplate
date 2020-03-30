package rooms

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jackc/pgx"
)

func (api *RoomsApi) allHandler() (rooms []Room, err error) {

	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM rooms`); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var room Room
		if err = rows.Scan(
			&room.Id,
			&room.HotelId,
			&room.Name,
			&room.Quantity,
			&room.Price,
		); err != nil {
			log.Println("allHandler(): ", err.Error())
			if err == pgx.ErrNoRows {
				rooms, err = []Room{}, nil
				return
			}
			return
		}
		rooms = append(rooms, room)
	}
	return
}

func (api *RoomsApi) addHandler(room Room) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO rooms (name, quantity, price, hotel_id) VALUES($1, $2, $3, $4) RETURNING id`,
		room.Name, room.Quantity, room.Price, room.HotelId).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}

func (api *RoomsApi) detailHandler(id int) (room Room, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM rooms WHERE id=$1`, id).Scan(
		&room.Id,
		&room.HotelId,
		&room.Name,
		&room.Quantity,
		&room.Price,
	); err != nil {
		log.Println("detailHandler(): ", err.Error())
		if err == pgx.ErrNoRows {
			room, err = Room{}, nil
			return
		}
		return
	}
	return
}

func (api *RoomsApi) updateHandler(room Room) (err error) {
	var result pgx.CommandTag
	result, err = api.Db.Exec(`
		UPDATE rooms SET name=$1, quantity=$2, price=$3, hotel_id=$4 WHERE id=$5`,
		room.Name,
		room.Quantity,
		room.Price,
		room.HotelId,
		room.Id,
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

func (api *RoomsApi) deleteHandler(id int) (err error) {
	var result pgx.CommandTag
	if result, err = api.Db.Exec(`DELETE FROM rooms WHERE id=$1`, id); err != nil {
		log.Println("deleteHandler(): ", err.Error())
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-delete")
		return
	}
	return
}

func (api *RoomsApi) bodyToStruct(body io.ReadCloser) (model Room, err error) {
	if err = json.NewDecoder(body).Decode(&model); err != nil {
		return
	}
	return
}
