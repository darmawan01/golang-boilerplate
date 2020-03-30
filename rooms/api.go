package rooms

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomsApi) Register() {
	api.Router.Handle("/rooms", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/rooms", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/rooms/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/rooms/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/rooms/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomsApi registered")
}

func (api *RoomsApi) all(w http.ResponseWriter, r *http.Request) {
	rooms, err := api.allHandler()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", rooms, 0, 0, 0),
	)
}

func (api *RoomsApi) add(w http.ResponseWriter, r *http.Request) {
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

func (api *RoomsApi) detail(w http.ResponseWriter, r *http.Request) {

	roomId := utils.GetIDParam(r)
	if utils.IsEmpty(roomId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required hotel_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of hotel_id is invalid", nil),
		)
	}

	room, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", room, 0, 0, 0),
	)

}

func (api *RoomsApi) update(w http.ResponseWriter, r *http.Request) {
	roomId := utils.GetIDParam(r)
	if utils.IsEmpty(roomId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required hotel_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of hotel_id is invalid", nil),
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

	if !utils.IsEmpty(updated.Name) {
		room.Name = updated.Name
	}

	if updated.Quantity != room.Quantity {
		room.Quantity = updated.Quantity
	}

	if updated.Price != room.Price {
		room.Price = updated.Price
	}

	if updated.HotelId != room.HotelId {
		room.HotelId = updated.HotelId
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

func (api *RoomsApi) delete(w http.ResponseWriter, r *http.Request) {
	roomId := utils.GetIDParam(r)
	if utils.IsEmpty(roomId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required hotel_id as param", nil),
		)
	}

	id, err := strconv.Atoi(roomId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of hotel_id is invalid", nil),
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
