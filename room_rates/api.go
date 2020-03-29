package room_rates

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomRateSApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomRateSApi) Register() {
	api.Router.Handle("/room-rates", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/room-rates", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/room-rates/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/room-rates/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/room-rates/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomRateSApi registered")
}

func (api *RoomRateSApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomRateSApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomRateSApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomRateSApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *RoomRateSApi) delete(w http.ResponseWriter, r *http.Request) {

}
