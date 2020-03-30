package room_rates

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomRatesApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomRatesApi) Register() {
	api.Router.Handle("/room-rates", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/room-rates", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/room-rates/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/room-rates/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/room-rates/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomRatesApi registered")
}

func (api *RoomRatesApi) all(w http.ResponseWriter, r *http.Request) {
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

func (api *RoomRatesApi) add(w http.ResponseWriter, r *http.Request) {
	room, err := api.bodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Invalid body", nil),
		)
		return
	}

	if ok, err := utils.ValidateStruct(room); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil))
		return

	}

	lastInsertedId, err := api.addHandler(room)
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

func (api *RoomRatesApi) detail(w http.ResponseWriter, r *http.Request) {

	roomRatesId := utils.GetIDParam(r)
	if utils.IsEmpty(roomRatesId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required room_rates_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomRatesId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of room_rates_id is invalid", nil),
		)
	}

	room, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", room, 0, 0, 0),
	)

}

func (api *RoomRatesApi) update(w http.ResponseWriter, r *http.Request) {
	roomRatesId := utils.GetIDParam(r)
	if utils.IsEmpty(roomRatesId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required room_rates_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomRatesId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of room_rates_id is invalid", nil),
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

	room, err := api.detailHandler(id)

	if !utils.IsEmpty(updated.Date.String()) {
		room.Date = updated.Date
	}

	if updated.Price != room.Price {
		room.Price = updated.Price
	}

	if updated.RoomId != room.RoomId {
		room.RoomId = updated.RoomId
	}

	if err = api.updateHandler(room); err != nil {
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

func (api *RoomRatesApi) delete(w http.ResponseWriter, r *http.Request) {
	roomRatesId := utils.GetIDParam(r)
	if utils.IsEmpty(roomRatesId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required room_rates_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomRatesId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of room_rates_id is invalid", nil),
		)
	}

	if err = api.deleteHandler(id); err != nil {
		if !utils.IsEqual(err.Error(), "failed-to-delete") {
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
