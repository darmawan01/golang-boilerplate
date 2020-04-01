package orders

import (
	"encoding/json"
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

	api.Router.Handle("/order/reports", http.HandlerFunc(api.reports)).Methods("GET")

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
		utils.DataFormat("Success !", roomRates),
	)
}

func (api *OrdersApi) add(w http.ResponseWriter, r *http.Request) {
	var order OrderRequests

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
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
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", order),
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
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

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

func (api *OrdersApi) reports(w http.ResponseWriter, r *http.Request) {
	reqParams := []string{"month", "year", "hotel", "group_by"}
	param := utils.GetQueryParam(r, reqParams...)

	report, err := api.reportsHandler(param)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", report),
	)
}
