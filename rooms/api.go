package rooms

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomsApi) Register() {
	api.Router.Handle("/order-guests", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/order-guests", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomsApi registered")
}

func (api *RoomsApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomsApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomsApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomsApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomsApi) delete(w http.ResponseWriter, r *http.Request) {

}
