package order_guests

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type OrderGuestSApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *OrderGuestSApi) Register() {
	api.Router.Handle("/order-guests", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/order-guests", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("OrderGuestSApi registered")
}

func (api *OrderGuestSApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderGuestSApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderGuestSApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderGuestSApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *OrderGuestSApi) delete(w http.ResponseWriter, r *http.Request) {

}
