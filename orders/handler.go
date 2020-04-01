package orders

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	orderGuests "kodingworks/order_guests"
	orderItems "kodingworks/order_items"

	"github.com/jackc/pgx"
)

var rows *pgx.Rows

func (api *OrdersApi) allHandler() (orders []Order, err error) {
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

func (api *OrdersApi) addHandler(order OrderRequests) (lastInsertedId int, err error) {
	if err = api.Db.QueryRow(`
		INSERT INTO orders (hotel_id, guest_id, status, check_in, check_out) 
		VALUES($1, $2, $3, $4, $5) RETURNING id`,
		order.HotelId,
		order.GuestId,
		order.Status,
		order.CheckinAt,
		order.CheckoutAt,
	).Scan(&lastInsertedId); err != nil {
		log.Println("addHandler(): ", err.Error())
		return
	}

	order.OrderGuests.OrderId = lastInsertedId
	orderGuestsApi := orderGuests.OrderGuestsApi{Db: api.Db}
	if _, err = orderGuestsApi.AddHandler(order.OrderGuests); err != nil {
		log.Println("orderGuestsApi.AddHandler(): ", err.Error())
		return
	}

	for _, item := range order.OrderItems {
		item.OrderId = lastInsertedId
		orderItemsApi := orderItems.OrderItemsApi{Db: api.Db}
		if _, err = orderItemsApi.AddHandler(item); err != nil {
			log.Println("orderItemsApi.AddHandler(): ", err.Error())
			return
		}
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
		UPDATE orders SET hotel_id=$1, guest_id=$2, status=$3, check_in=$4, 
		check_out=$5 WHERE id=$6`,
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

func (api *OrdersApi) reportsHandler(param map[string]string) (rep report, err error) {

	query := `
		SELECT item.price, DATE(o.created_at) AS date FROM orders o 
		JOIN (SELECT price, order_id FROM order_items) AS item ON item.order_id=o.id 
		JOIN (SELECT id FROM hotels) AS h ON h.id=o.hotel_id
	`
	query += " WHERE "

	if len(param["month"]) > 0 {
		query += fmt.Sprintf("AND DATE_PART('month', o.created_at)='%s' ", param["month"])
	}
	if len(param["year"]) > 0 {
		query += fmt.Sprintf("AND DATE_PART('year', o.created_at)='%s' ", param["year"])
	}
	if len(param["hotel"]) > 0 {
		query += fmt.Sprintf("AND h.id='%s' ", param["hotel"])
	}
	if len(param["group_by"]) > 0 {
		query = strings.Replace(query, "SELECT item.price", "SELECT SUM(item.price)", 1)
		query += " GROUP BY DATE(o.created_at)"
	}

	// Clean query
	query = strings.Replace(query, "WHERE AND", "WHERE", 1)

	if rows, err = api.Db.Query(query); err != nil {
		log.Println("reportsHandler(): ", err.Error())
		return
	}
	defer rows.Close()

	totals := map[string]interface{}{}
	datePrev, pricePrev := "", 0.0
	for rows.Next() {
		var price float64
		var date time.Time
		if err = rows.Scan(
			&price,
			&date,
		); err != nil && err != pgx.ErrNoRows {
			log.Println("rows.Scan(): ", err.Error())
			return
		}
		rep.TotalSales += price
		dateStr := fmt.Sprintf("%s", date.Local().Format("01-02-2006"))
		if datePrev != dateStr {
			rep.TotalOrders++
		} else if datePrev == dateStr {
			price += pricePrev
		}
		totals[dateStr] = map[string]interface{}{
			"price":  fmt.Sprintf("%.2f", price),
			"orders": fmt.Sprintf("%v", rep.TotalOrders+1),
		}
		datePrev, pricePrev = dateStr, price
	}
	rep.Totals = totals
	return
}
