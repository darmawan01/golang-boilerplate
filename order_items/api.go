package order_items

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type OrderItemsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *OrderItemsApi) Register() {
	api.Router.Handle("/order-items/{id}", http.HandlerFunc(api.detail)).Methods("GET")

	log.Println("OrderItemsApi registered")
}

func (api *OrderItemsApi) detail(w http.ResponseWriter, r *http.Request) {

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

	order, err := api.getByOrder(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", order, 0, 0, 0),
	)

}
