package rooms

import "github.com/jackc/pgx"

var rows *pgx.Rows

func (api *RoomsApi) listHandler() (rooms []Room, err error) {

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
