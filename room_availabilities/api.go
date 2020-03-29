package room_availabalities

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomAvailabilitiesApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomAvailabilitiesApi) Register() {
	api.Router.Handle("/room-availabilies", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/room-availabilies", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/room-availabilies/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/room-availabilies/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/room-availabilies/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomAvailabilitiesApi registered")
}

func (api *RoomAvailabilitiesApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomAvailabilitiesApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomAvailabilitiesApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomAvailabilitiesApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomAvailabilitiesApi) delete(w http.ResponseWriter, r *http.Request) {

}
