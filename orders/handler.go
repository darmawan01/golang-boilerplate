package orders

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jackc/pgx"
)

func (api *OrdersApi) allHandler() (orders []Order, err error) {
	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM orders`); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		if err = rows.Scan(
			&order.Id,
			&order.HotelId,
			&order.GuestId,
			&order.Status,
			&order.CheckinAt,
			&order.CheckoutAt,
			&order.CreatedAt,
		); err != nil {
			log.Println("allHandler(): ", err.Error())
			if err == pgx.ErrNoRows {
				orders, err = []Order{}, nil
				return
			}
			return
		}
		orders = append(orders, order)
	}
	return
}

func (api *OrdersApi) addHandler(order Order) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO orders (hotel_id, guest_id, status, check_in, check_out) 
		VALUES($1, $2, $3, $4, $5, %6) RETURNING id`,
		order.HotelId,
		order.GuestId,
		order.Status,
		order.CheckinAt,
		order.CheckoutAt,
	).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}

func (api *OrdersApi) detailHandler(id int) (order Order, err error) {
	if err = api.Db.QueryRow(`SELECT * FROM orders WHERE id=$1`, id).Scan(
		&order.Id,
		&order.HotelId,
		&order.GuestId,
		&order.Status,
		&order.CheckinAt,
		&order.CheckoutAt,
		&order.CreatedAt,
	); err != nil {
		log.Println("detailHandler(): ", err.Error())
		if err == pgx.ErrNoRows {
			order, err = Order{}, nil
			return
		}
		return
	}
	return
}

func (api *OrdersApi) updateHandler(order Order) (err error) {
	var result pgx.CommandTag
	result, err = api.Db.Exec(`
		UPDATE orders SET hotel_id=$1, guest_id=$2, status=$3, check_in=%4, 
		check_out=%5 WHERE id=$6`,
		order.HotelId,
		order.GuestId,
		order.Status,
		order.CheckinAt,
		order.CheckoutAt,
		order.Id,
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

func (api *OrdersApi) bodyToStruct(body io.ReadCloser) (model Order, err error) {
	if err = json.NewDecoder(body).Decode(&model); err != nil {
		return
	}
	return
}
