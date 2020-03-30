package order_guests

import (
	"log"

	"github.com/jackc/pgx"
)

func (api *OrderGuestSApi) getByOrder(orderId int) (orderGuests []OrderGuest, err error) {

	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM order_guests WHERE order_id=$1`, orderId); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem OrderGuest
		if err = rows.Scan(
			&orderItem.Id,
			&orderItem.OrderId,
			&orderItem.Name,
			&orderItem.Email,
			&orderItem.PhoneNumber,
		); err != nil {
			log.Println("getByOrder(): ", err.Error())
			if err == pgx.ErrNoRows {
				orderGuests, err = []OrderGuest{}, nil
				return
			}
			return
		}
		orderGuests = append(orderGuests, orderItem)
	}
	return
}

func (api *OrderGuestSApi) addHandler(orderItem OrderGuest) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO order_items (order_id, name, email, phone_number) VALUES($1, $2, $3, $4) RETURNING id`,
		orderItem.OrderId, orderItem.Name, orderItem.Email).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}
