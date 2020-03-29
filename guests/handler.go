package guests

import (
	"fmt"

	"github.com/jackc/pgx"
)

func (api *GuestsApi) allHandler() (guests []Guest, err error) {
	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM guests`); err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var guest Guest

		if err = rows.Scan(&guest.Id, &guest.Name, &guest.Email, &guest.PhoneNumber, &guest.IdentificationId); err != nil {
			if err == pgx.ErrNoRows {
				guests, err = []Guest{}, nil
				return
			}
			return
		}
		guests = append(guests, guest)
	}
	return
}

func (api *GuestsApi) addHandler(guest Guest) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		CREATE guests (name, email, phone_number, identification_id) VALUES($1, $2, $3, $4) RETURNING id`,
		guest.Name, guest.Email, guest.PhoneNumber, guest.IdentificationId).Scan(&lastInsertedId); err != nil {
		return
	}
	return
}

func (api *GuestsApi) detailHandler(id int) (guest Guest, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM guests WHERE id=$1`, id).Scan(
		&guest.Id, &guest.Name, &guest.Email, &guest.PhoneNumber, &guest.IdentificationId); err != nil {
		if err == pgx.ErrNoRows {
			guest, err = Guest{}, nil
			return
		}
		return
	}
	return
}

func (api *GuestsApi) updateHandler(guest Guest) (err error) {
	var result pgx.CommandTag
	result, err = api.Db.Exec(`
		UPDATE guests SET name=$1, email=$2, phone_number=$3, identification_id=$4 WHERE id=$1`,
		guest.Name,
		guest.Email,
		guest.PhoneNumber,
		guest.IdentificationId,
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

func (api *GuestsApi) deleteHandler(id int) (err error) {
	var result pgx.CommandTag
	if result, err = api.Db.Exec(`DELETE FROM guests WHERE id=$1`); err != nil {
		return
	}
	if result.RowsAffected() == 0 {
		err = fmt.Errorf("failed-to-delete")
		return
	}
	return
}
