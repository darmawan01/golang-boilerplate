package orders

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type OrdersApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *OrdersApi) Register() {
	api.Router.Handle("/orders", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/orders", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/orders/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/orders/{id}", http.HandlerFunc(api.update)).Methods("PUT")

	log.Println("OrdersApi registered")
}

func (api *OrdersApi) all(w http.ResponseWriter, r *http.Request) {
	roomRates, err := api.allHandler()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", roomRates, 0, 0, 0),
	)
}

func (api *OrdersApi) add(w http.ResponseWriter, r *http.Request) {
	order, err := api.bodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Invalid body", nil),
		)
		return
	}

	if ok, err := utils.ValidateStruct(order); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil))
		return

	}

	lastInsertedId, err := api.addHandler(order)
	if err != nil {
		log.Println(err)
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return

	}
	utils.RespondwithJSON(
		w,
		http.StatusCreated,
		map[string]interface{}{
			"lastInsertedId": lastInsertedId,
		},
	)

}

func (api *OrdersApi) detail(w http.ResponseWriter, r *http.Request) {

	orderId := utils.GetIDParam(r)
	if utils.IsEmpty(orderId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required id as param", nil),
		)
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of id is invalid", nil),
		)
	}

	order, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", order, 0, 0, 0),
	)

}

func (api *OrdersApi) update(w http.ResponseWriter, r *http.Request) {
	orderId := utils.GetIDParam(r)
	if utils.IsEmpty(orderId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required id as param", nil),
		)
	}

	id, err := strconv.Atoi(orderId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of id is invalid", nil),
		)
	}

	updated, err := api.bodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Invalid body", nil),
		)
		return
	}

	if ok, err := utils.ValidateStruct(updated); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return
	}

	order, err := api.detailHandler(id)

	if updated.HotelId != order.HotelId {
		order.HotelId = updated.HotelId
	}

	if updated.GuestId != order.GuestId {
		order.GuestId = updated.GuestId
	}

	if updated.Status != order.Status {
		order.Status = updated.Status
	}

	if !utils.IsEmpty(updated.CheckinAt.String()) {
		order.CheckinAt = updated.CheckinAt
	}

	if !utils.IsEmpty(updated.CheckoutAt.String()) {
		order.CheckoutAt = updated.CheckoutAt
	}

	if err = api.updateHandler(order); err != nil {
		if !utils.IsEqual(err.Error(), "failed-to-update") {
			utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
			return
		}
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, map[string]interface{}{
		"messages": "Success !",
	})
}
