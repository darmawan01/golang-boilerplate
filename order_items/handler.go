package order_items

import (
	"log"

	"github.com/jackc/pgx"
)

func (api *OrderItemsApi) getByOrder(orderId int) (orderItems []OrderItem, err error) {

	var rows *pgx.Rows
	if rows, err = api.Db.Query(`SELECT * FROM order_items WHERE order_id=$1`, orderId); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem OrderItem
		if err = rows.Scan(
			&orderItem.Id,
			&orderItem.OrderId,
			&orderItem.RoomId,
			&orderItem.Quantity,
			&orderItem.Price,
		); err != nil {
			log.Println("getByOrder(): ", err.Error())
			if err == pgx.ErrNoRows {
				orderItems, err = []OrderItem{}, nil
				return
			}
			return
		}
		orderItems = append(orderItems, orderItem)
	}
	return
}

func (api *OrderItemsApi) addHandler(orderItem OrderItem) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO order_items (order_id, room_id, quantity, price) VALUES($1, $2, $3, $4) RETURNING id`,
		orderItem.OrderId, orderItem.RoomId, orderItem.Quantity).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}
	return
}
