package room_rates

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jackc/pgx"
)

func (api *RoomRatesApi) allHandler() (roomRates []RoomRate, err error) {
	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM room_rates`); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var roomRate RoomRate
		if err = rows.Scan(
			&roomRate.Id,
			&roomRate.RoomId,
			&roomRate.Date,
			&roomRate.Price,
		); err != nil {
			log.Println("allHandler(): ", err.Error())
			if err == pgx.ErrNoRows {
				roomRates, err = []RoomRate{}, nil
				return
			}
			return
		}
		roomRates = append(roomRates, roomRate)
	}
	return
}

func (api *RoomRatesApi) addHandler(roomRate RoomRate) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO room_rates (date, price, room_id) VALUES($1, $2, $3) RETURNING id`,
		roomRate.Date, roomRate.Price, roomRate.RoomId).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}

func (api *RoomRatesApi) detailHandler(id int) (roomRate RoomRate, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM room_rates WHERE id=$1`, id).Scan(
		&roomRate.Id,
		&roomRate.RoomId,
		&roomRate.Date,
		&roomRate.Price,
	); err != nil {
		log.Println("detailHandler(): ", err.Error())
		if err == pgx.ErrNoRows {
			roomRate, err = RoomRate{}, nil
			return
		}
		return
	}
	return
}

func (api *RoomRatesApi) updateHandler(roomRate RoomRate) (err error) {
	var result pgx.CommandTag
	result, err = api.Db.Exec(`
		UPDATE room_rates SET date=$1, price=$2, room_id=$3 WHERE id=$4`,
		roomRate.Date,
		roomRate.Price,
		roomRate.RoomId,
		roomRate.Id,
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

func (api *RoomRatesApi) deleteHandler(id int) (err error) {
	var result pgx.CommandTag
	if result, err = api.Db.Exec(`DELETE FROM room_rates WHERE id=$1`, id); err != nil {
		log.Println("deleteHandler(): ", err.Error())
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-delete")
		return
	}
	return
}

func (api *RoomRatesApi) bodyToStruct(body io.ReadCloser) (model RoomRate, err error) {
	if err = json.NewDecoder(body).Decode(&model); err != nil {
		return
	}
	return
}
