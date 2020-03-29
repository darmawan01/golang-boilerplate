package hotels

import (
	"fmt"

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
		CREATE hotels (name, address, latitude, longitude) VALUES($1, $2, $3, $4) RETURNING id`,
		hotel.Name, hotel.Address, hotel.Latitute, hotel.Longitude).Scan(&lastInsertedId); err != nil {
		return
	}
	return
}

func (api *HotelsApi) detailHandler(id int) (hotel Hotel, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM hotels WHERE id=$1`, id).Scan(
		&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Latitute, &hotel.Longitude); err != nil {
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
		UPDATE hotels SET name=$1, address=$2, latitude=$3, longitude=$4 WHERE id=$1`,
		hotel.Name,
		hotel.Address,
		hotel.Latitute,
		hotel.Longitude,
	)
	if err != nil {
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
	if result, err = api.Db.Exec(`DELETE FROM hotels WHERE id=$1`); err != nil {
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-delete")
		return
	}
	return
}
