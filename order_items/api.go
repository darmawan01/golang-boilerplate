package order_items

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type OrderItemSApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *OrderItemSApi) Register() {
	api.Router.Handle("/order-items", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/order-items", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/order-items/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/order-items/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/order-items/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("OrderItemSApi registered")
}

func (api *OrderItemSApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderItemSApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderItemSApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderItemSApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderItemSApi) delete(w http.ResponseWriter, r *http.Request) {

}
