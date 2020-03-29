package guests

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type GuestsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *GuestsApi) Register() {
	api.Router.Handle("/guests", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/guests", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/guests/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/guests/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/guests/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("GuestsApi registered")
}

func (api *GuestsApi) all(w http.ResponseWriter, r *http.Request) {

}

func (api *GuestsApi) add(w http.ResponseWriter, r *http.Request) {

}

func (api *GuestsApi) detail(w http.ResponseWriter, r *http.Request) {

}

func (api *GuestsApi) update(w http.ResponseWriter, r *http.Request) {

}

func (api *GuestsApi) delete(w http.ResponseWriter, r *http.Request) {

}
