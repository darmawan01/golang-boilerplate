package room_availabalities

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomAvailabilitiesApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomAvailabilitiesApi) Register() {
	api.Router.Handle("/room-availability", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/room-availability", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/room-availability/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/room-availability/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/room-availability/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomAvailabilitiesApi registered")
}

func (api *RoomAvailabilitiesApi) all(w http.ResponseWriter, r *http.Request) {
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

func (api *RoomAvailabilitiesApi) add(w http.ResponseWriter, r *http.Request) {
	roomAvailability, err := api.bodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Invalid body", nil),
		)
		return
	}

	if ok, err := utils.ValidateStruct(roomAvailability); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil))
		return

	}

	lastInsertedId, err := api.addHandler(roomAvailability)
	if err != nil {
		if strings.Contains(err.Error(), "insert or update on table") {
			utils.RespondwithJSON(w, http.StatusBadRequest,
				utils.ErrFormat("Room with room_id not found", nil),
			)
			return
		}
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

func (api *RoomAvailabilitiesApi) detail(w http.ResponseWriter, r *http.Request) {

	roomAvailabilityId := utils.GetIDParam(r)
	if utils.IsEmpty(roomAvailabilityId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required room_availability_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomAvailabilityId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of room_availability_id is invalid", nil),
		)
	}

	roomAvailability, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", roomAvailability),
	)

}

func (api *RoomAvailabilitiesApi) update(w http.ResponseWriter, r *http.Request) {
	roomAvailabilityId := utils.GetIDParam(r)
	if utils.IsEmpty(roomAvailabilityId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required room_availability_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomAvailabilityId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of room_availability_id is invalid", nil),
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

	roomAvailability, err := api.detailHandler(id)

	if !utils.IsEmpty(updated.Date.String()) {
		roomAvailability.Date = updated.Date
	}

	if updated.Quantity != roomAvailability.Quantity {
		roomAvailability.Quantity = updated.Quantity
	}

	if updated.RoomId != roomAvailability.RoomId {
		roomAvailability.RoomId = updated.RoomId
	}

	if err = api.updateHandler(roomAvailability); err != nil {
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

func (api *RoomAvailabilitiesApi) delete(w http.ResponseWriter, r *http.Request) {
	roomAvailabilityId := utils.GetIDParam(r)
	if utils.IsEmpty(roomAvailabilityId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required room_availability_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomAvailabilityId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of room_availability_id is invalid", nil),
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
