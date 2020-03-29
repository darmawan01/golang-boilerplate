package orders

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type OrdersApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *OrdersApi) Register() {
	api.Router.Handle("/order-guests", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/order-guests", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("OrdersApi registered")
}

func (api *OrdersApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *OrdersApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *OrdersApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *OrdersApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *OrdersApi) delete(w http.ResponseWriter, r *http.Request) {

}